# Way

Procedurally generated battle-kart racing game built in Go with Ebitengine.

## Visual Perspective

Way uses a mandatory **over-the-shoulder** gameplay perspective for all racing gameplay. This camera view provides an immersive third-person perspective positioned behind and slightly above the player's kart, offering optimal visibility of the track ahead while maintaining awareness of the kart's position and orientation.

## Directory Structure

```
way/
├── cmd/
│   ├── client/          # Game client (Ebitengine window)
│   └── server/          # Dedicated race server
├── pkg/
│   ├── engine/          # ECS world, components, and game systems
│   ├── procgen/         # Procedural generation interface and seed PRNG
│   ├── rendering/       # Runtime sprite and render system
│   ├── audio/           # Procedural audio system
│   ├── network/         # Multiplayer networking system
│   └── config/          # Viper configuration loading
├── config.yaml          # Default configuration file
├── go.mod
└── ROADMAP.md
```

## Build

```sh
go build ./cmd/client
go build ./cmd/server
```

## Run

```sh
./client   # opens game window
./server   # starts dedicated server
```

## Configuration

Configuration is loaded from `config.yaml` in the working directory or `$HOME/.way/config.yaml`. All values have sensible defaults and the application runs without a config file present.

The camera perspective is mandated to be "over-the-shoulder" and cannot be changed. Camera distance, height, and angle can be adjusted in the configuration file.

See `config.yaml` for available options.

## Dependencies

- [Ebitengine v2](https://ebitengine.org/) — 2D game engine
- [Viper](https://github.com/spf13/viper) — configuration management
