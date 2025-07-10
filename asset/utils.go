package asset

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

func imageFromBytes(pngBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func flipImageXAxis(image *ebiten.Image) *ebiten.Image {
	width, height := image.Bounds().Size().X, image.Bounds().Size().Y
	flipped := ebiten.NewImage(width, height)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(width), 0)
	image.DrawImage(flipped, op)

	return flipped
}
