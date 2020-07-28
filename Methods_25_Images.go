package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
