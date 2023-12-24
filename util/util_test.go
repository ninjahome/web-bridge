package util

import (
	"fmt"
	"testing"
)

func TestTxt2Img(t *testing.T) {
	var err = ConvertLongTweetToImg("我爱我的足球")
	fmt.Println(err)
}
