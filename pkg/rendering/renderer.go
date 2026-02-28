package rendering

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/opd-ai/way/pkg/engine"
)

// RenderSystem generates and draws runtime sprites, tiles, and particles.
type RenderSystem struct {
	Screen *ebiten.Image
}

func (s *RenderSystem) Update(w *engine.World, dt float64) {
	_ = w
	_ = dt
}

// Draw renders the current frame to the screen.
func (s *RenderSystem) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 30, A: 255})
}
