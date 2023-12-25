package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

import (
	"errors"
)

const (
	MaxReqContentLen = 1024 * 1024 * 5
	MaxTwitterLen    = 280
)

var (
	ErrSignInvalid      = errors.New("invalid signature")
	ErrSignNoAddr       = errors.New("no public address found")
	ErrSignNotMatch     = errors.New("signature address not match")
	ErrHttpEmptyRequest = errors.New("empty http post")
)

func RandomBytesInHex(count int) string {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(buf)
}

func toJSON(v interface{}) string {
	js, _ := json.Marshal(v)
	return string(js)
}

func ParseTemplates(path string) *template.Template {
	fs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var files []string
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".html") {
			files = append(files, filepath.Join(path, f.Name()))
		}
	}
	tmpl := template.New("").Funcs(template.FuncMap{"json": toJSON})
	return template.Must(tmpl.ParseFiles(files...))
}

func Verify(address, message, signedMessage string) error {
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
	addressOrig := strings.ToLower(address)
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

func ReadRequest(request *http.Request, obj any) error {
	body := io.LimitReader(request.Body, MaxReqContentLen)
	var b bytes.Buffer
	n, err := b.ReadFrom(body)
	if err != nil && n != 0 {
		return err
	}
	if n == 0 {
		return ErrHttpEmptyRequest
	}
	return json.Unmarshal(b.Bytes(), obj)
}

func IsOverTwitterLimit(text string) bool {
	return utf8.RuneCountInString(text) > MaxTwitterLen
}
func TruncateString(str string, n int) string {
	if n > utf8.RuneCountInString(str) {
		return str
	}
	return string([]rune(str)[:n])
}
