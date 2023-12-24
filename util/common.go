package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/freetype"
	"html/template"
	"image"
	"image/draw"
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
	TweetFontInImg   = "Noto_Sans_SC.ttf"
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

func ConvertLongTweetToImg(txt string) (image.Image, error) {
	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	// 读取字体
	fontBytes, err := os.ReadFile("util/Noto_Sans_SC.ttf")
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	// 设置字体参数
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(24)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(24)>>6))
	_, err = c.DrawString(txt, pt)
	if err != nil {
		return nil, err
	}
	return img, nil

}
