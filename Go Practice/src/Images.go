package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	weight int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.weight, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
