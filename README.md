# Way

Procedural battle-kart racing game built with Go and Ebitengine.

## Directory Structure

```
cmd/
  client/          Client binary (Ebitengine window)
  server/          Dedicated race server binary
pkg/
  engine/          ECS world, components, systems, PCG system
  procgen/         Procedural generation framework and generators
  rendering/       Runtime sprite and tile generation
  audio/           Procedural audio synthesis
  network/         Authoritative server, client prediction, protocol
config/
  config.go        Viper-based configuration loading
config.yaml        Default configuration file
```

## Build

```
go build ./...
```

## Run

Client:

```
go run ./cmd/client
```

Server:

```
go run ./cmd/server
```

## Configuration

The game reads `config.yaml` from the working directory. If the file is missing, built-in defaults are used. Settings can also be overridden via environment variables.

See `config.yaml` for available options (window size, game seed, genre, server address, audio, debug flags).

## Dependencies

- [Go](https://go.dev/) 1.24+
- [Ebitengine](https://ebitengine.org/) v2 — game engine
- [Viper](https://github.com/spf13/viper) — configuration management
