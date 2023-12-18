package server

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

var (
	ErrSignInvalid  = errors.New("invalid signature")
	ErrSignNoAddr   = errors.New("no public address found")
	ErrSignNotMatch = errors.New("signature address not match")
)

type SignInParam struct {
	Address   string `json:"address"`
	SignTim   uint64 `json:"signTim"`
	Signature string `json:"signature"`
}

func (sp *SignInParam) Verify(message, signedMessage string) error {
	sig := common.FromHex(signedMessage)

	// 在以太坊中，签名是65字节，最后一个字节是V值（27或28），但在Geth内部需要将其减少到0或1
	if sig[64] != 27 && sig[64] != 28 {
		return ErrSignInvalid
	}
	sig[64] -= 27

	// 使用恢复方法来获取公钥
	pubKey, err := crypto.SigToPub(signHash(message), sig)
	if err != nil {
		return ErrSignNoAddr
	}

	// 将公钥转换为以太坊地址
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	addressOrig := strings.ToLower(sp.Address)
	addressSigned := strings.ToLower(recoveredAddr.Hex())
	if addressOrig != addressSigned {
		return ErrSignNotMatch
	}
	return nil
}

func signHash(data string) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
