package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func GenerateRandomData(passphraseFlag string) (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := privateKey.D.Bytes()
	privateKeyInt := privateKey.D

	if len(privateKeyInt.Bytes()) > 32 {
		return "", "", fmt.Errorf("private key too long for uint256")
	}

	harsher := sha3.NewLegacyKeccak256()
	harsher.Write(privateKeyBytes)
	hash := harsher.Sum(nil)
	var hashStr = fmt.Sprintf("0x%x", hash)

	harsher = sha256.New()
	harsher.Write([]byte(passphraseFlag))
	key := harsher.Sum(nil)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", "", err
	}

	encryptedData := gcm.Seal(nonce, nonce, []byte(privateKeyInt.String()), nil)
	return hex.EncodeToString(encryptedData), hashStr, nil
}

func DecryptRandomData(encryptedData string, passphrase string) (string, error) {
	data, err := hex.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	harsher := sha256.New()
	harsher.Write([]byte(passphrase))
	key := harsher.Sum(nil)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}
