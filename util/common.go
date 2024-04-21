package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rivo/uniseg"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

import (
	"errors"
)

const (
	MaxReqContentLen = 1 << 23
	MaxTwitterLen    = 280
	urlWeight        = 23
)

var (
	Version             string
	Commit              string
	BuildTime           string
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

var urlRegex = regexp.MustCompile(`https?://\S+`)

// ParseTweet calculates the length of a tweet text based on Twitter's specific rules.
func ParseTweet(text string) (int, bool) {
	weightedLength := 0
	urls := urlRegex.FindAllStringIndex(text, -1)

	// Subtract urls and add fixed url length
	lastIndex := 0
	for _, loc := range urls {
		// Add the text before the URL
		weightedLength += calculateWeight(text[lastIndex:loc[0]])
		// Add fixed URL length
		weightedLength += urlWeight
		lastIndex = loc[1]
	}
	// Add the rest of the text after the last URL
	weightedLength += calculateWeight(text[lastIndex:])

	// Check if the tweet is valid based on its weighted length
	isValid := weightedLength <= MaxTwitterLen
	return weightedLength, isValid
}

// calculateWeight calculates the weight of the text.
func calculateWeight(text string) int {
	weight := 0
	gr := uniseg.NewGraphemes(text)
	for gr.Next() {
		r := gr.Runes()
		if isCJK(r) {
			weight += 2 // CJK characters have weight 2
		} else {
			weight += 1 // Non-CJK characters have weight 1
		}
	}
	return weight
}

// isCJK checks if the rune slice is a CJK character.
func isCJK(runes []rune) bool {
	for _, r := range runes {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func IsValidTwitterContent(text string) bool {
	characterCount, valid := ParseTweet(text)
	LogInst().Debug().Int("text-len", len(text)).Str("text", text).Int("rune-len", characterCount).Msg("tweet text length")
	return valid
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
