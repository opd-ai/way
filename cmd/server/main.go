// Package main implements the Way dedicated race server.
package main

import (
	"fmt"
	"log"

	"github.com/opd-ai/way/config"
	"github.com/opd-ai/way/pkg/network"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv := network.NewServer(
		cfg.Server.Address,
		cfg.Server.Port,
		cfg.Server.TickRate,
		cfg.Server.MaxPlayers,
	)

	fmt.Printf("Way server starting on %s:%d (tick rate: %d Hz, max players: %d)\n",
		cfg.Server.Address, cfg.Server.Port, cfg.Server.TickRate, cfg.Server.MaxPlayers)

	if err := srv.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
