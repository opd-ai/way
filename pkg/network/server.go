package network

import (
	"github.com/opd-ai/way/pkg/engine"
)

// Server is the authoritative race server.
type Server struct {
	Address    string
	Port       int
	TickRate   int
	MaxPlayers int
	World      *engine.World
}

// NewServer creates a Server with the given settings.
func NewServer(address string, port, tickRate, maxPlayers int) *Server {
	return &Server{
		Address:    address,
		Port:       port,
		TickRate:   tickRate,
		MaxPlayers: maxPlayers,
		World:      engine.NewWorld(),
	}
}

// Start begins listening for client connections and ticking the game world.
func (s *Server) Start() error {
	// Skeleton: server networking will be implemented in Phase 4
	return nil
}

// Stop gracefully shuts down the server.
func (s *Server) Stop() error {
	// Skeleton: server shutdown will be implemented in Phase 4
	return nil
}
