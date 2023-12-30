package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKeyBytes := privateKey.D.Bytes()
	harsher := sha3.NewLegacyKeccak256()
	harsher.Write(privateKeyBytes)
	hash := harsher.Sum(nil)
	fmt.Printf("\n0x%x\n", hash)
}
