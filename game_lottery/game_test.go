package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/web-bridge/util"
	"log"
	"os"
	"testing"
)

var (
	passphraseFlag = ""
	data           = ""
)

func init() {
	flag.StringVar(&passphraseFlag, "pwd", "", "--pwd=[password]")
	flag.StringVar(&data, "data", "", "--data=[DATA]")
}

func createAccount(password string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	keyStore := keystore.NewKeyStore(".", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := keyStore.ImportECDSA(privateKey, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet created: ", account.Address.Hex())
}
func testData() {
	decryptFlag := flag.Bool("d", false, "Decrypt mode")
	encryptedDataFlag := flag.String("e", "", "Encrypted data for decryption")

	flag.Parse()

	bts, err := hex.DecodeString(passphraseFlag)

	if err != nil {
		panic(err)
	}
	if *decryptFlag == false {
		fmt.Println(util.GenerateRandomData(bts))
		return
	}

	if passphraseFlag == "" || *encryptedDataFlag == "" {
		fmt.Println("Passphrase and encrypted data are required for decryption")
		flag.Usage()
		return
	}
	fmt.Println(util.DecryptRandomData(*encryptedDataFlag, bts))
	return
}

// go test -run TestNewAccount --pwd=
func TestNewAccount(t *testing.T) {
	createAccount(passphraseFlag)
}

// go test -run TestGenerateEncryptedRandomHash --pwd= --data=
func TestGenerateEncryptedRandomHash(t *testing.T) {

	jsonBytes, err := os.ReadFile("dessage.key")
	if err != nil {
		panic(err)
	}

	key, err := keystore.DecryptKey(jsonBytes, passphraseFlag)
	if err != nil {
		t.Fatal(err)
	}

	block, err := aes.NewCipher(key.PrivateKey.D.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		t.Fatal(err)
	}

	encryptedData := gcm.Seal(nonce, nonce, []byte(data), nil)
	fmt.Println(hex.EncodeToString(encryptedData))

	decrypted, err := util.DecryptRandomData(hex.EncodeToString(encryptedData), key.PrivateKey.D.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypted.String())
}

func TestInfuraInfo(t *testing.T) {
	client, err := ethclient.Dial("https://arbitrum-mainnet.infura.io/v3/08db2487445e45fe848b3b7b6b95c080")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve block number: %v", err)
	}

	fmt.Println("Block number:", blockNumber.Time())
}
