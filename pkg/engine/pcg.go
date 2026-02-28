package engine

import (
	"github.com/opd-ai/way/pkg/procgen"
)

// PCGSystem drives all procedural generators.
type PCGSystem struct {
	Seed   int64
	Genre  procgen.GenreID
	Params procgen.GenerationParams
}

// NewPCGSystem creates a PCGSystem with the given seed and genre.
func NewPCGSystem(seed int64, genre procgen.GenreID) *PCGSystem {
	return &PCGSystem{
		Seed:  seed,
		Genre: genre,
		Params: procgen.GenerationParams{
			Genre: genre,
			Seed:  seed,
		},
	}
}

// Update satisfies the System interface. PCGSystem is typically run once at
// world initialisation rather than every tick.
func (p *PCGSystem) Update(w *World, dt float64) {
	// Skeleton: PCG orchestration will be implemented in Phase 1–2
}
