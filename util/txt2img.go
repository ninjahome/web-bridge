package util

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"os"
)

func calcTextHeight(txt string, maxWidth int, fontSize float64, face font.Face) (int, error) {
	// 用于测量的临时变量
	var width, height, lineHeight fixed.Int26_6
	lineHeight = fixed.I(int(fontSize * 1.5))

	for _, r := range txt {
		if r == '\n' {
			// 显式换行
			width = 0
			height += lineHeight
			continue
		}

		awidth, ok := face.GlyphAdvance(r)
		if !ok {
			continue
		}

		if width+awidth > fixed.I(maxWidth) {
			// 自动换行
			width = 0
			height += lineHeight
		}
		width += awidth
	}

	// 加上最后一行的高度
	height += lineHeight

	return height.Ceil(), nil
}

func ConvertLongTweetToImg(txt string) (image.Image, error) {
	// 读取字体
	fontBytes, err := os.ReadFile("Noto_Sans_SC.ttf")
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	// 创建字体Face
	opts := &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	}
	face := truetype.NewFace(f, opts)

	// 原本的最大宽度
	originalMaxWidth := 500
	// 为右边界留出空间
	rightPadding := 40
	maxWidth := originalMaxWidth - rightPadding

	textHeight, err := calcTextHeight(txt, maxWidth, 24, face)
	if err != nil {
		return nil, err
	}

	// 创建新的画布
	img := image.NewRGBA(image.Rect(0, 0, originalMaxWidth, textHeight))
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	// 设置 freetype 上下文参数
	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(24)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	// 绘制文本
	pt := freetype.Pt(0, int(c.PointToFixed(24)>>6))
	var width fixed.Int26_6
	lineHeight := fixed.I(int(24 * 1.5))

	for _, r := range txt {
		if r == '\n' {
			// 显式换行
			width = 0
			pt.Y += lineHeight
			pt.X = 0
			continue
		}

		awidth, ok := face.GlyphAdvance(r)
		if !ok {
			continue
		}

		if width+awidth > fixed.I(maxWidth) {
			// 自动换行
			width = 0
			pt.Y += lineHeight
			pt.X = 0
		}

		pt.X += awidth
		width += awidth

		// 绘制字符
		_, err = c.DrawString(string(r), pt)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}
