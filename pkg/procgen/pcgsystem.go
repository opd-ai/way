package procgen

import "github.com/opd-ai/way/pkg/engine"

// PCGSystem drives all generators via the Generator interface.
type PCGSystem struct {
	Generators []Generator
}

func (s *PCGSystem) Update(w *engine.World, dt float64) {
	_ = w
	_ = dt
}
