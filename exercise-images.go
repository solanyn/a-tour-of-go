package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct{
    w, h int
    c uint8
}

func (im *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, im.w, im.h)
}

func (im *Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img *Image) At(x, y int) color.Color {
    return color.RGBA{img.c + uint8(x*y + x*y), img.c + uint8(x*y + x*y), 255, 255}
}

func main() {
	m := Image{500, 500, 200}
	pic.ShowImage(&m)
}

