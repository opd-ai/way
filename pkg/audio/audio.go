package audio

import "github.com/opd-ai/way/pkg/engine"

// AudioSystem handles procedural engine sounds, Doppler, adaptive music, and SFX.
type AudioSystem struct{}

func (s *AudioSystem) Update(w *engine.World, dt float64) {
	_ = w
	_ = dt
}
