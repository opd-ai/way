package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opd-ai/way/pkg/config"
	"github.com/opd-ai/way/pkg/engine"
	"github.com/opd-ai/way/pkg/network"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	world := engine.NewWorld()
	world.AddSystem(&engine.PhysicsSystem{})
	world.AddSystem(&engine.RaceSystem{})
	world.AddSystem(&engine.ItemSystem{})
	world.AddSystem(&engine.TrackSystem{})
	world.AddSystem(&network.NetworkSystem{
		IsServer:   true,
		TickRateHz: cfg.Server.TickRateHz,
	})

	if cfg.Server.TickRateHz <= 0 {
		fmt.Fprintf(os.Stderr, "config: server.tick_rate_hz must be > 0, got %d\n", cfg.Server.TickRateHz)
		os.Exit(1)
	}

	log.Printf("starting server loop at %d Hz (configured address: %s)", cfg.Server.TickRateHz, cfg.Server.Address)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(time.Second / time.Duration(cfg.Server.TickRateHz))
	defer ticker.Stop()

	last := time.Now()
	for {
		select {
		case <-stop:
			log.Println("server shutting down")
			return
		case <-ticker.C:
			now := time.Now()
			dt := now.Sub(last).Seconds()
			last = now
			world.Tick(dt)
		}
	}
}
