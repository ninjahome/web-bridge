package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
)

func GenerateRandomData(passphrase []byte) (string, []byte, error) {
	randomNum, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", nil, err
	}
	randomNumBytes := randomNum.D.Bytes()

	if len(randomNumBytes) != 32 {
		return "", nil, fmt.Errorf("invalid random number")
	}

	harsher := sha3.NewLegacyKeccak256()
	harsher.Write(randomNumBytes)
	hash := harsher.Sum(nil)

	block, err := aes.NewCipher(passphrase)
	if err != nil {
		return "", nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", nil, err
	}

	encryptedData := gcm.Seal(nonce, nonce, []byte(randomNum.D.String()), nil)
	return hex.EncodeToString(encryptedData), hash, nil
}

func DecryptRandomData(encryptedData string, passphrase []byte) (*big.Int, error) {
	data, err := hex.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(passphrase)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	random, ok := big.NewInt(0).SetString(string(decryptedData), 10)
	if !ok {
		return nil, fmt.Errorf("invalid big int data")
	}
	return random, nil
}
