package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ninjahome/web-bridge/util"
	"testing"
)

var (
	passphraseFlag = ""
)

func init() {
	passphraseFlag = *flag.String("p", "", "Encryption passphrase")
}
func createAccount(password string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	keyStore := keystore.NewKeyStore("dessage.key", keystore.StandardScryptN, keystore.StandardScryptP)
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

func TestNewAccount(t *testing.T) {
	createAccount(passphraseFlag)
}

func TestNewConf(t *testing.T) {

	//bts, _ := json.Marshal(gs)

	//_ = os.WriteFile("game.conf", bts, 0644)
}
