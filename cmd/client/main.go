package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/opd-ai/way/pkg/audio"
	"github.com/opd-ai/way/pkg/config"
	"github.com/opd-ai/way/pkg/engine"
	"github.com/opd-ai/way/pkg/network"
	"github.com/opd-ai/way/pkg/procgen"
	"github.com/opd-ai/way/pkg/rendering"
)

// Game implements the ebiten.Game interface.
type Game struct {
	world    *engine.World
	renderer *rendering.RenderSystem
	lastTime time.Time
}

func (g *Game) Update() error {
	now := time.Now()
	dt := now.Sub(g.lastTime).Seconds()
	g.lastTime = now
	g.world.Tick(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

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
	world.AddSystem(&procgen.PCGSystem{})
	world.AddSystem(&audio.AudioSystem{})
	world.AddSystem(&network.NetworkSystem{})

	renderer := &rendering.RenderSystem{}
	world.AddSystem(renderer)

	game := &Game{
		world:    world,
		renderer: renderer,
		lastTime: time.Now(),
	}

	ebiten.SetWindowSize(cfg.Window.Width, cfg.Window.Height)
	ebiten.SetWindowTitle(cfg.Window.Title)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
