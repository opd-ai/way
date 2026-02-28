// Package rendering provides the runtime sprite and tile generation pipeline.
package rendering

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/opd-ai/way/pkg/engine"
)

// RenderSystem generates and draws all visual content at runtime.
type RenderSystem struct {
	screenWidth  int
	screenHeight int
}

// NewRenderSystem creates a RenderSystem with the given screen dimensions.
func NewRenderSystem(width, height int) *RenderSystem {
	return &RenderSystem{
		screenWidth:  width,
		screenHeight: height,
	}
}

// Update satisfies the engine.System interface.
func (r *RenderSystem) Update(w *engine.World, dt float64) {
	// Skeleton: rendering logic will be implemented in Phase 2
}

// Draw renders all entities to the screen.
func (r *RenderSystem) Draw(screen *ebiten.Image, w *engine.World) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 30, A: 255})
	// Skeleton: sprite drawing will be implemented in Phase 2
}
