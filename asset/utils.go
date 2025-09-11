package asset

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func imageFromBytes(pngBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(pngBytes))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func flipImageXAxis(image *ebiten.Image) *ebiten.Image {
	width, height := image.Bounds().Dx(), image.Bounds().Dy()
	flipped := ebiten.NewImage(width, height)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(width), 0)
	flipped.DrawImage(image, op)
	return flipped
}

func rotateImageClockwise(image *ebiten.Image) *ebiten.Image {
	// get the original image dimensions
	originWidth := image.Bounds().Dx()
	originHeight := image.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}

	// make the center of the image the origin, so we can rotate around this
	op.GeoM.Translate(-float64(originWidth)/2, -float64(originHeight)/2)

	// rotate the matrix 90 degrees
	op.GeoM.Rotate(math.Pi / 2)

	// move the origin back to 0,0
	op.GeoM.Translate(float64(originHeight)/2, float64(originWidth)/2) // +200, +50

	// render the rotated image and return it
	rotatedWidth := originHeight
	rotatedHeight := originWidth
	rotated := ebiten.NewImage(rotatedWidth, rotatedHeight)
	rotated.DrawImage(image, op)

	return rotated
}
