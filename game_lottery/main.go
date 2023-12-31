package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/web-bridge/blockchain/ethapi"
	"github.com/ninjahome/web-bridge/database"
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

func initConfig(filePath string) *server.SysConf {
	cf := new(server.SysConf)

	bts, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bts, &cf); err != nil {
		panic(err)
	}
	util.SetLogLevel(cf.LogLevel)
	fmt.Println(cf.String())
	return cf
}

func readWallet(filePath string) *keystore.Key {
	for {
		fmt.Print("Enter Password: ")
		passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("\nError reading password:" + err.Error())
			continue
		}

		password := string(passwordBytes)

		jsonBytes, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		key, err := keystore.DecryptKey(jsonBytes, password)
		if err != nil {
			fmt.Println("failed to decrypt wallet:", err.Error())
			continue
		}
		fmt.Println("open wallet success:", key.Address.String())
		return key
	}
}
func main() {

	walletFile := flag.String("wallet", "dessage.key", "wallet file")
	confFile := flag.String("conf", "config.json", "config file ")
	startRoundRandom := flag.String("random", "", "start round random number")
	startRoundNo := flag.Int("round-no", 0, "start round no")
	version := flag.Bool("version", false, "game_lottery --version")
	flag.Parse()

	if *version {
		fmt.Println("\n==================================================")
		fmt.Printf("Version:\t%s\n", util.Version)
		fmt.Printf("Build:\t\t%s\n", util.BuildTime)
		fmt.Printf("Commit:\t\t%s\n", util.Commit)
		fmt.Println("==================================================")
		return
	}
	cf := initConfig(*confFile)
	key := readWallet(*walletFile)

	gs := NewGame(key, cf)
	if len(*startRoundRandom) > 0 {
		if err := gs.SetupFirstRound(*startRoundRandom, *startRoundNo); err != nil {
			panic(err)
		}
	}

	go gs.Server()

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
	key         *keystore.Key
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
	DiscoverTxHash  string `json:"discover_hash"  firestore:"discover_hash"`
	LastRoundStatus int8   `json:"last_round_status"  firestore:"last_round_status"`
}

type GameResult struct {
	RoundNo string `json:"round_no"`
	Random  string `json:"random"`
	Success bool   `json:"success"`
}

func NewGame(key *keystore.Key, cf *server.SysConf) *GameService {
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
	return &GameService{key: key,
		conf:     cf,
		checker:  t,
		fileCli:  client,
		ctx:      ctx,
		txResult: make(chan *GameResult, 1)}
}

func (gs *GameService) performGameCheck() (*big.Int, bool) {
	util.LogInst().Info().Msg("time to check game status")
	if gs.isWaitingTx {
		util.LogInst().Info().Msg("game checking:waiting for transaction packaging")
		return nil, false
	}
	curNo, nextTime, err := gs.gameTimeOn()
	if err != nil {
		util.LogInst().Err(err).Msg("check game status failed")
		return nil, false
	}

	ti := time.Now().Add(time.Duration(gs.conf.GameTimeInMinute) * time.Minute)
	if ti.Sub(*nextTime) <= 0 {
		util.LogInst().Debug().Str("current-round", curNo.String()).
			Str("next-time", nextTime.String()).Msg("time is not on")
		return curNo, false
	}
	util.LogInst().Info().Int64("round-no", curNo.Int64()).Msg("start to find winner")
	return curNo, true
}

func (gs *GameService) Server() {
	util.LogInst().Info().Msg("game server start.......")
	for {
		select {
		case <-gs.checker.C:
			curNo, ok := gs.performGameCheck()
			if !ok {
				util.LogInst().Info().Msg("game still in progress")
				continue
			}

			curRandomNumber, err := gs.loadCurrentEncryptedRandom(curNo)
			if err != nil {
				util.LogInst().Err(err).Str("current-round", curNo.String()).Msg("failed to load random raw data")
				continue
			}
			privateBytes := gs.key.PrivateKey.D.Bytes()
			nextRandomEncrypted, nextHash, err := util.GenerateRandomData(privateBytes)
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
				DiscoverTxHash:  tx,
				LastRoundStatus: TxStatusInit,
			}
			err = gs.saveDiscoverInfo(curNo, meta)
			if err != nil {
				util.LogInst().Err(err).Msg("save game meta failed")
			}

		case result := <-gs.txResult:
			util.LogInst().Debug().Bool("status", result.Success).
				Str("current-round", result.RoundNo).
				Msg("transaction package success")

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
		util.LogInst().Err(err).Msg("dial eth failed")
		return nil, nil, err
	}
	defer cli.Close()

	contractAddress := common.HexToAddress(gs.conf.GameContract)
	game, err := ethapi.NewTweetLotteryGame(contractAddress, cli)
	if err != nil {
		util.LogInst().Err(err).Str("contract-address", gs.conf.GameContract).Msg("failed create game obj")
		return nil, nil, err
	}

	roundNo, err := game.CurrentRoundNo(nil)
	if err != nil {
		util.LogInst().Err(err).Str("contract-address", gs.conf.GameContract).Msg("failed to fetch current round no from blockchain")
		return nil, nil, err
	}

	result, err := game.GameInfoRecord(nil, roundNo)
	if err != nil {
		util.LogInst().Err(err).Str("current-round", roundNo.String()).
			Msg("failed to fetch game info of current round")
		return nil, nil, err
	}

	discoverTime := time.Unix(result.DiscoverTime.Int64(), 0)
	bts, _ := json.Marshal(result)
	util.LogInst().Debug().Str("current-round", roundNo.String()).
		Str("discover-time", discoverTime.String()).
		Msg(string(bts))
	return roundNo, &discoverTime, nil
}

func (gs *GameService) getTxClient() (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(gs.conf.InfuraUrl)
	if err != nil {
		util.LogInst().Err(err).Msg("dial to block chain api point failed")
		return nil, nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(gs.key.PrivateKey, big.NewInt(gs.conf.ChainID))
	if err != nil {
		util.LogInst().Err(err).Msg("suggest gas failed")
		return nil, nil, err
	}

	return client, auth, nil
}

func (gs *GameService) discoverWinner(random *big.Int, nextRoundRandomHash []byte, curNo *big.Int) (string, error) {
	util.LogInst().Info().Str("current-no", curNo.String()).Msg("start to discover winner")
	client, auth, err := gs.getTxClient()
	if err != nil {
		util.LogInst().Err(err).Msg("get transaction client failed")
		return "", err
	}
	defer client.Close()

	contractAddress := common.HexToAddress(gs.conf.GameContract)
	game, err := ethapi.NewTweetLotteryGame(contractAddress, client)
	if err != nil {
		util.LogInst().Err(err).Str("contract-address", gs.conf.GameContract).
			Msg("failed to create game from contract")
		return "", err
	}
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
		util.LogInst().Debug().Str("current-no", curNo.String()).
			Str("tx-hash", tx.Hash().String()).
			Msg("check receipt status")
		select {
		case <-queryTicker.C:
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				util.LogInst().Err(err).Str("current-no", curNo.String()).
					Str("tx-hash", tx.Hash().String()).
					Msg("Error checking transaction receipt:")
				continue
			}

			if receipt == nil {
				util.LogInst().Info().Str("current-no", curNo.String()).
					Str("tx-hash", tx.Hash().String()).Msg("transaction is in process")
				continue
			}

			result.Success = receipt.Status == types.ReceiptStatusSuccessful

			util.LogInst().Warn().Str("current-no", curNo.String()).
				Str("tx-hash", tx.Hash().String()).Bool("tx-status", result.Success).
				Msg("winner discover transaction finished")
			gs.txResult <- &result
			return
		}
	}
}

func (gs *GameService) loadCurrentEncryptedRandom(no *big.Int) (*big.Int, error) {
	opCtx, cancel := context.WithTimeout(gs.ctx, database.DefaultDBTimeOut)
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

	privateKeyBytes := gs.key.PrivateKey.D.Bytes()
	curRandomNumber, err := util.DecryptRandomData(gr.EncryptedRandom, privateKeyBytes)
	if err != nil {
		util.LogInst().Err(err).Msg("decrypt random of current round from encrypted data failed")
		return nil, err
	}
	util.LogInst().Debug().Str("current-round", no.String()).Msg("random raw data load success")
	return curRandomNumber, nil
}

func (gs *GameService) saveDiscoverInfo(no *big.Int, nextRandom GameRandomMeta) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, database.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(no.Add(no, big.NewInt(1)).String())
	_, err := randomDoc.Set(opCtx, nextRandom)
	return err
}

func (gs *GameService) SetupFirstRound(startRandom string, roundNo int) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, database.DefaultDBTimeOut)
	defer cancel()
	randomDoc := gs.fileCli.Collection(DBTableGameRandom).Doc(big.NewInt(int64(roundNo)).String())
	nextRandom := GameRandomMeta{
		EncryptedRandom: startRandom,
		LastRoundStatus: TxStatusSuccess,
	}
	_, err := randomDoc.Set(opCtx, nextRandom)
	return err
}

func (gs *GameService) updateDiscoverInfo(result *GameResult) error {
	opCtx, cancel := context.WithTimeout(gs.ctx, database.DefaultDBTimeOut)
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
