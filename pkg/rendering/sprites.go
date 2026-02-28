package rendering

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// GenerateSprite creates a procedural sprite image with the given dimensions and base color.
func GenerateSprite(width, height int, base color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	img.Fill(base)
	return img
}
