package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ninjahome/web-bridge/util"
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

func TestTweetLength(t *testing.T) {
	text := "【第十三章】宠辱若惊，贵大患若身。何谓宠辱若惊？宠为下，得之若惊，失之若惊何患？故贵以身为天下，若可寄天下\\\\r\\\\nCurrent Prize Pool: 0.264ETH，Vote on this tweet and participate in the prize pool. link:https://sharp-happy-grouse.ngrok-free.app/buyRights?NjTID=1713690380862"
	length, valid := util.ParseTweet(text)
	fmt.Printf("Calculated Twitter length: %d\n", length)
	fmt.Printf("Calculated Twitter valid: %t\n", valid)
}
