package identico

import (
	"image"
	"image/color"
	"image/draw"
)

func Classic(mask image.Image, bg, fg color.Color) image.Image {
	bounds := mask.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	bgimg := FillBackground(w, h, bg)
	fgimg := ReplaceMask(mask, fg)

	dst := image.NewNRGBA(bounds)
	draw.Draw(dst, bounds, bgimg, image.ZP, draw.Src)
	draw.Draw(dst, bounds, fgimg, image.ZP, draw.Over)
	return dst
}

func FillBackground(width, height int, col color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.ZP, draw.Src)
	return img
}

func ReplaceMask(mask image.Image, col color.Color) image.Image {
	bounds := mask.Bounds()
	dst := image.NewNRGBA(bounds)
	w, h := bounds.Max.X, bounds.Max.Y
	r, g, b, _ := col.RGBA()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pixel := mask.At(x, y)
			_, _, _, alpha := pixel.RGBA()
			if alpha != 0 {
				rgba := color.NRGBA{shift(r), shift(g), shift(b), shift(alpha)}
				dst.Set(x, y, rgba)
			} else {
				dst.Set(x, y, pixel)
			}
		}
	}
	return dst
}

func shift(v uint32) uint8 {
	return uint8(v >> 8)
}
