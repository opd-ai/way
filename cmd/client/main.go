// Package main implements the Way game client.
package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/opd-ai/way/config"
	"github.com/opd-ai/way/pkg/audio"
	"github.com/opd-ai/way/pkg/engine"
	"github.com/opd-ai/way/pkg/rendering"
)

// Game implements the ebiten.Game interface.
type Game struct {
	world    *engine.World
	renderer *rendering.RenderSystem
	audio    *audio.AudioSystem
	cfg      *config.Config
}

// NewGame creates a Game instance with the given configuration.
func NewGame(cfg *config.Config) *Game {
	w := engine.NewWorld()

	renderer := rendering.NewRenderSystem(cfg.Window.Width, cfg.Window.Height)
	audioSys := audio.NewAudioSystem(cfg.Audio.Enabled, cfg.Audio.Volume)

	w.RegisterSystem(renderer)
	w.RegisterSystem(audioSys)
	w.RegisterSystem(&engine.PhysicsSystem{})
	w.RegisterSystem(&engine.RaceSystem{})
	w.RegisterSystem(&engine.ItemSystem{})
	w.RegisterSystem(&engine.TrackSystem{})

	return &Game{
		world:    w,
		renderer: renderer,
		audio:    audioSys,
		cfg:      cfg,
	}
}

// Update advances the game state by one tick.
func (g *Game) Update() error {
	const dt = 1.0 / 60.0
	g.world.Tick(dt)
	return nil
}

// Draw renders the game to the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen, g.world)
}

// Layout returns the game's logical screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.cfg.Window.Width, g.cfg.Window.Height
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Way — genre: %s, seed: %d\n", cfg.Game.Genre, cfg.Game.Seed)

	game := NewGame(cfg)

	ebiten.SetWindowSize(cfg.Window.Width, cfg.Window.Height)
	ebiten.SetWindowTitle(cfg.Window.Title)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
