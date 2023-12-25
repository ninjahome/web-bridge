package util

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
)

func calcTextHeight(txt string, maxWidth int, fontSize float64, face font.Face) (int, error) {
	var _, height, lineHeight fixed.Int26_6
	lineHeight = fixed.I(int(fontSize * 1.5))

	lines := splitTextIntoLines(txt, maxWidth, face)
	for range lines {
		height += lineHeight
	}

	height += lineHeight

	return height.Ceil(), nil
}

func splitTextIntoLines(txt string, maxWidth int, face font.Face) []string {
	var lines []string
	var line string
	var width fixed.Int26_6
	for _, r := range txt {
		if r == '\n' {
			lines = append(lines, line)
			line = ""
			width = 0
			continue
		}

		aWidth, ok := face.GlyphAdvance(r)
		if !ok {
			continue
		}

		if width+aWidth > fixed.I(maxWidth) {
			lines = append(lines, line)
			line = ""
			width = 0
		}
		line += string(r)
		width += aWidth
	}
	lines = append(lines, line) // Add last line
	return lines
}

func ConvertLongTweetToImg(txt string, f *truetype.Font, fontSize float64) (image.Image, error) {

	opts := &truetype.Options{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingNone,
	}
	face := truetype.NewFace(f, opts)

	originalMaxWidth := 1600
	leftPadding := 40
	maxWidth := originalMaxWidth - 2*leftPadding

	textHeight, err := calcTextHeight(txt, maxWidth, fontSize, face)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, originalMaxWidth, textHeight))
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	c := freetype.NewContext()
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	pt := freetype.Pt(leftPadding, int(c.PointToFixed(fontSize)>>6))
	lineHeight := fixed.I(int(fontSize * 1.5))

	lines := splitTextIntoLines(txt, maxWidth, face)
	for _, line := range lines {
		_, err = c.DrawString(line, pt)
		if err != nil {
			return nil, err
		}
		pt.Y += lineHeight
	}

	return img, nil
}
