package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rivo/uniseg"
	"html/template"
	"io"
	"math/big"
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

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	// It's negative.
	if number.IsInt64() {
		return rpc.BlockNumber(number.Int64()).String()
	}
	// It's negative and large, which is invalid.
	return fmt.Sprintf("<invalid %d>", number)
}

func GetBlockByNumber(url string, blockNum *big.Int) (*Block, error) {
	noStr := toBlockNumArg(blockNum)
	payload := bytes.NewBuffer([]byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["%s", false],"id":1}`, noStr)))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		LogInst().Err(err).Msg("Error creating request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		LogInst().Err(err).Msg("Error sending request to server")
		return nil, err
	}
	defer resp.Body.Close()
	var response JsonResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		LogInst().Err(err).Msg("decode response failed")
		return nil, err
	}
	response.Result.TimeStamp2 = hexutil.MustDecodeBig(response.Result.Timestamp)
	return &response.Result, nil
}

type Block struct {
	BaseFeePerGas    string `json:"baseFeePerGas"`
	Difficulty       string `json:"difficulty"`
	ExtraData        string `json:"extraData"`
	GasLimit         string `json:"gasLimit"`
	GasUsed          string `json:"gasUsed"`
	Hash             string `json:"hash"`
	L1BlockNumber    string `json:"l1BlockNumber"`
	LogsBloom        string `json:"logsBloom"`
	Miner            string `json:"miner"`
	MixHash          string `json:"mixHash"`
	Nonce            string `json:"nonce"`
	Number           string `json:"number"`
	ParentHash       string `json:"parentHash"`
	ReceiptsRoot     string `json:"receiptsRoot"`
	SendCount        string `json:"sendCount"`
	SendRoot         string `json:"sendRoot"`
	Sha3Uncles       string `json:"sha3Uncles"`
	Size             string `json:"size"`
	StateRoot        string `json:"stateRoot"`
	Timestamp        string `json:"timestamp"`
	TimeStamp2       *big.Int
	TotalDifficulty  string   `json:"totalDifficulty"`
	Transactions     []string `json:"transactions"`
	TransactionsRoot string   `json:"transactionsRoot"`
	Uncles           []string `json:"uncles"`
}

type JsonResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Block  `json:"result"`
}
