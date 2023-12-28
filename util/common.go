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

func Verify(address, message, signedMessage string) (string, error) {
	sig := common.FromHex(signedMessage)

	if sig[64] != 27 && sig[64] != 28 {
		return "", ErrSignInvalid
	}
	sig[64] -= 27

	var prefixedHash = signHash(message)
	pubKey, err := crypto.SigToPub(prefixedHash, sig)
	if err != nil {
		return "", ErrSignNoAddr
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	addressOrig := strings.ToLower(address)
	addressSigned := strings.ToLower(recoveredAddr.Hex())
	if addressOrig != addressSigned {
		return "", ErrSignNotMatch
	}
	return common.BytesToHash(prefixedHash).Hex(), nil
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
	return len(text) > MaxTwitterLen
}

func TruncateString(raw, append string) string {
	appendLen := len(append) + 3 // 加上三个省略号的长度
	if appendLen >= MaxTwitterLen {
		return append[:MaxTwitterLen] // 如果附加文本太长，只返回截断的附加文本
	}

	maxLen := MaxTwitterLen - appendLen
	truncated := ""
	for _, r := range raw {
		if len(truncated)+len(string(r)) > maxLen {
			break
		}
		truncated += string(r)
	}

	return truncated + "..." + append
}
