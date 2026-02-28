// Package audio provides procedural audio generation: engine sounds,
// adaptive music, SFX, and Doppler effects.
package audio

import (
	"github.com/opd-ai/way/pkg/engine"
)

// AudioSystem synthesises all game audio at runtime.
type AudioSystem struct {
	Enabled bool
	Volume  float64
}

// NewAudioSystem creates an AudioSystem with the given settings.
func NewAudioSystem(enabled bool, volume float64) *AudioSystem {
	return &AudioSystem{
		Enabled: enabled,
		Volume:  volume,
	}
}

// Update satisfies the engine.System interface.
func (a *AudioSystem) Update(w *engine.World, dt float64) {
	if !a.Enabled {
		return
	}
	// Skeleton: audio tick logic will be implemented in Phase 3
}
