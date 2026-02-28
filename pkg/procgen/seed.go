package procgen

import "math/rand/v2"

// DeterministicRNG wraps a seeded PRNG to ensure reproducible output.
type DeterministicRNG struct {
	*rand.Rand
}

// NewDeterministicRNG creates a new deterministic PRNG from a seed.
func NewDeterministicRNG(seed int64) *DeterministicRNG {
	return &DeterministicRNG{
		Rand: rand.New(rand.NewPCG(uint64(seed), 0)),
	}
}
