package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/web-bridge/blockchain/sol/ethapi"
	"github.com/ninjahome/web-bridge/server"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/crypto/ssh/terminal"
	"google.golang.org/api/option"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	DBTableGameRandom = "game_lottery_random"
)
const (
	TxStatusInit = iota
	TxStatusSuccess
	TxStatusFailed
)

func main() {
	walletFile := flag.String("wallet", "dessage.key", "wallet file")
	confFile := flag.String("conf", "game.conf", "config file ")
	firstRoundRandom := flag.String("random", "", "first round random number")
	cf := new(server.SysConf)

	bts, err := os.ReadFile(*confFile)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bts, &cf); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	fmt.Print("Enter Password: ")
	passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("\nError reading password")
		return
	}

	password := string(passwordBytes)

	jsonBytes, err := os.ReadFile(*walletFile)
	if err != nil {
		panic(err)
	}

	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		panic(err)
	}

	gs := NewGame(key.PrivateKey, cf)
	go gs.Server()
	if len(*firstRoundRandom) > 0 {
		if err := gs.SetupFirstRound(*firstRoundRandom); err != nil {
			panic(err)
		}
	}

	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>service start at pid(%s)<<<<<<<<<<\n", pid)
	if err := os.WriteFile("gs.pid", []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)
	sig := <-sigCh
	fmt.Printf("\n>>>>>>>>>>service finished(%s)<<<<<<<<<<\n", sig)
}

type GameService struct {
	privateKey  *ecdsa.PrivateKey
	conf        *server.SysConf
	checker     *time.Ticker
	fileCli     *firestore.Client
	ctx         context.Context
	isWaitingTx bool
	txResult    chan *GameResult
}

type GameRandomMeta struct {
	EncryptedRandom string `json:"encrypted_random"  firestore:"encrypted_random"`
	RandomHash      string `json:"random_hash"  firestore:"random_hash"`
	RandomLastRound string `json:"random_last_round"  firestore:"random_last_round"`
	DiscoverHash    string `json:"discover_hash"  firestore:"discover_hash"`
	LastRoundStatus int8   `json:"last_round_status"  firestore:"last_round_status"`
}

type GameResult struct {
	RoundNo string `json:"round_no"`
	Random  string `json:"random"`
	Success bool   `json:"success"`
}

func NewGame(key *ecdsa.PrivateKey, cf *server.SysConf) *GameService {
	util.SetLogLevel(cf.LogLevel)
	ctx := context.Background()
	var client *firestore.Client
	var err error
	if cf.LocalRun {
		_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
		client, err = firestore.NewClient(ctx, cf.ProjectID)
	} else {
		client, err = firestore.NewClient(ctx, cf.ProjectID, option.WithCredentialsFile(cf.KeyFilePath))
	}
	if err != nil {
		panic(err)
	}
	t := time.NewTicker(time.Duration(cf.GameTimeInMinute) * time.Minute)
	return &GameService{privateKey: key, conf: cf, checker: t, fileCli: client, ctx: ctx, txResult: make(chan *GameResult, 1)}
}

func (gs *GameService) performGameCheck() (*big.Int, bool) {
	if gs.isWaitingTx {
		return nil, false
	}
	curNo, nextTime, err := gs.gameTimeOn()
	if err != nil {
		util.LogInst().Err(err).Msg("check game status failed")
		return nil, false
	}

	util.LogInst().Info().Msg("start to check game time")

	ti := time.Now().Add(10 * time.Minute)
	if ti.Sub(*nextTime) <= 0 {
		return curNo, false
	}
	util.LogInst().Info().Int64("round-no", curNo.Int64()).Msg("start to find winner")
	return curNo, true
}

func (gs *GameService) Server() {

	for {
		select {
		case <-gs.checker.C:
			curNo, ok := gs.performGameCheck()
			if !ok {
				util.LogInst().Info().Msg("game still in progress")
				continue
			}

			curRandomNumber, _ := gs.loadCurrentEncryptedRandom(curNo)

			nextRandomEncrypted, nextHash, err := util.GenerateRandomData(gs.privateKey.D.Bytes())
			if err != nil {
				util.LogInst().Err(err).Msg("generate random for next round failed")
				continue
			}

			tx, err := gs.discoverWinner(curRandomNumber, nextHash, curNo)
			if err != nil {
				util.LogInst().Err(err).Msg("discover winner failed")
				continue
			}

			var meta = GameRandomMeta{
				RandomHash:      hex.EncodeToString(nextHash),
				EncryptedRandom: nextRandomEncrypted,
				DiscoverHash:    tx,
				LastRoundStatus: TxStatusInit,
			}
			err = gs.saveDiscoverInfo(curNo, meta)
			if err != nil {
				util.LogInst().Err(err).Msg("save game meta failed")
			}

		case result := <-gs.txResult:
			err := gs.updateDiscoverInfo(result)
			if err != nil {
				util.LogInst().Err(err).Msg("update discover result failed")
			}
		}
	}
}

func (gs *GameService) gameTimeOn() (*big.Int, *time.Time, error) {
	cli, err := ethclient.Dial(gs.conf.InfuraUrl)
	if err != nil {
		return nil, nil, err
	}
	defer cli.Close()

	contractAddress := common.HexToAddress(gs.conf.GameContract)
	game, err := ethapi.NewTweetLotteryGame(contractAddress, cli)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}
	roundNo, err := game.CurrentRoundNo(nil)

	if err != nil {
		return nil, nil, err
	}
	result, err := game.GameInfoRecord(nil, roundNo)
	if err != nil {
		return nil, nil, err
	}

	discoverTime := time.Unix(result.DiscoverTime.Int64(), 0)
	return roundNo, &discoverTime, nil
}

func (gs *GameService) getTxClient() (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(gs.conf.InfuraUrl)
	if err != nil {
		util.LogInst().Err(err).Msg("dial to block chain api point failed")
		return nil, nil, err
	}

	publicKey := gs.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		util.LogInst().Warn().Msg("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		util.LogInst().Err(err).Msg("pending nonce failed")
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		util.LogInst().Err(err).Msg("suggest gas failed")
		return nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(gs.privateKey, big.NewInt(gs.conf.ChainID))
	if err != nil {
		util.LogInst().Err(err).Msg("suggest gas failed")
		return nil, nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return client, auth, nil
}

func (gs *GameService) discoverWinner(random *big.Int, nextRoundRandomHash []byte, curNo *big.Int) (string, error) {
	client, auth, err := gs.getTxClient()
	if err != nil {
		util.LogInst().Err(err).Msg("get transaction client failed")
		return "", err
	}

	contractAddress := common.HexToAddress(gs.conf.GameContract)
	game, err := ethapi.NewTweetLotteryGame(contractAddress, client)
	var byteArray [32]byte
	copy(byteArray[:32], nextRoundRandomHash)

	tx, err := game.DiscoverWinner(auth, random, byteArray)
	if err != nil {
		util.LogInst().Err(err).Msg("suggest gas failed")
		return "", err
	}

	util.LogInst().Info().Str("tx", tx.Hash().String()).Msg("init transaction success")

	gs.isWaitingTx = true
	go gs.waitTransactionResult(tx, client, random, curNo)
	return tx.Hash().String(), nil
}

func (gs *GameService) waitTransactionResult(tx *types.Transaction, client *ethclient.Client, random, curNo *big.Int) {
	queryTicker := time.NewTicker(time.Second * time.Duration(gs.conf.TxCheckerInSeconds))
	defer func() {
		gs.isWaitingTx = false
		queryTicker.Stop()
	}()
	var result = GameResult{
		RoundNo: curNo.String(),
		Random:  random.String(),
	}
	for {
		select {
		case <-queryTicker.C:
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				util.LogInst().Err(err).Msg("Error checking transaction receipt:")
				continue
			}

			if receipt == nil {
				util.LogInst().Info().Msg("transaction is in process")
				continue
			}

			if receipt.Status != types.ReceiptStatusSuccessful {
				util.LogInst().Warn().Msg("transaction failed, block number:")
				result.Success = false
				gs.txResult <- &result
			}

			util.LogInst().Warn().Msg("winner discover success")
			result.Success = true
			gs.txResult <- &result
		}
	}
}

func (gs *GameService) loadCurrentEncryptedRandom(no *big.Int) (*big.Int, error) {
	opCtx, cancel := context.WithTimeout(gs.ctx, server.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(no.String())
	doc, err := randomDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Msg("get random of lottery game failed")
		return nil, err
	}
	var gr GameRandomMeta
	err = doc.DataTo(&gr)
	if err != nil {
		util.LogInst().Err(err).Msg("parse random data of lottery game failed")
		return nil, err
	}

	privateKeyBytes := gs.privateKey.D.Bytes()
	curRandomNumber, err := util.DecryptRandomData(gr.EncryptedRandom, privateKeyBytes)
	if err != nil {
		util.LogInst().Err(err).Msg("decrypt random of current round from encrypted data failed")
		return nil, err
	}
	return curRandomNumber, nil
}

func (gs *GameService) saveDiscoverInfo(no *big.Int, nextRandom GameRandomMeta) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, server.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(no.Add(no, big.NewInt(1)).String())
	_, err := randomDoc.Set(opCtx, nextRandom)
	return err
}

func (gs *GameService) SetupFirstRound(s string) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, server.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(big.NewInt(0).String())
	nextRandom := GameRandomMeta{
		EncryptedRandom: s,
	}
	_, err := randomDoc.Set(opCtx, nextRandom)
	return err
}

func (gs *GameService) updateDiscoverInfo(result *GameResult) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, server.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(result.RoundNo)
	var resultStatus = TxStatusFailed
	if result.Success {
		resultStatus = TxStatusSuccess
	}
	_, err := randomDoc.Update(opCtx, []firestore.Update{
		{Path: "random_last_round", Value: result.Random},
		{Path: "last_round_status", Value: resultStatus},
	})

	return err
}
