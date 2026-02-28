package network

import "github.com/opd-ai/way/pkg/engine"

// NetworkSystem handles authoritative server, client prediction, and reconciliation.
type NetworkSystem struct {
	IsServer   bool
	TickRateHz int
}

func (s *NetworkSystem) Update(w *engine.World, dt float64) {
	_ = w
	_ = dt
}
