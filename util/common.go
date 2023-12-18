package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func RandomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", fmt.Errorf("could not generate %d random bytes: %v", count, err)
	}

	return hex.EncodeToString(buf), nil
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

	return template.Must(template.ParseFiles(files...))
}
