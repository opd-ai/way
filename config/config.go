// Package config provides Viper-based configuration loading for Way.
package config

import (
	"github.com/spf13/viper"
)

// Config holds all configuration values for the game.
type Config struct {
	Window WindowConfig `mapstructure:"window"`
	Game   GameConfig   `mapstructure:"game"`
	Server ServerConfig `mapstructure:"server"`
	Audio  AudioConfig  `mapstructure:"audio"`
	Debug  DebugConfig  `mapstructure:"debug"`
}

// WindowConfig holds display settings.
type WindowConfig struct {
	Width      int    `mapstructure:"width"`
	Height     int    `mapstructure:"height"`
	Title      string `mapstructure:"title"`
	Fullscreen bool   `mapstructure:"fullscreen"`
}

// GameConfig holds core game settings.
type GameConfig struct {
	TickRate int    `mapstructure:"tick_rate"`
	Seed     int64  `mapstructure:"seed"`
	Genre    string `mapstructure:"genre"`
}

// ServerConfig holds network server settings.
type ServerConfig struct {
	Address    string `mapstructure:"address"`
	Port       int    `mapstructure:"port"`
	TickRate   int    `mapstructure:"tick_rate"`
	MaxPlayers int    `mapstructure:"max_players"`
}

// AudioConfig holds audio settings.
type AudioConfig struct {
	Enabled bool    `mapstructure:"enabled"`
	Volume  float64 `mapstructure:"volume"`
}

// DebugConfig holds debug/development settings.
type DebugConfig struct {
	ShowFPS      bool `mapstructure:"show_fps"`
	ShowHitboxes bool `mapstructure:"show_hitboxes"`
}

// Load reads configuration from config.yaml and environment variables.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Defaults
	viper.SetDefault("window.width", 800)
	viper.SetDefault("window.height", 600)
	viper.SetDefault("window.title", "Way")
	viper.SetDefault("window.fullscreen", false)
	viper.SetDefault("game.tick_rate", 60)
	viper.SetDefault("game.seed", 0)
	viper.SetDefault("game.genre", "fantasy")
	viper.SetDefault("server.address", "localhost")
	viper.SetDefault("server.port", 7777)
	viper.SetDefault("server.tick_rate", 20)
	viper.SetDefault("server.max_players", 8)
	viper.SetDefault("audio.enabled", true)
	viper.SetDefault("audio.volume", 0.8)
	viper.SetDefault("debug.show_fps", true)
	viper.SetDefault("debug.show_hitboxes", false)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		// Config file not found; use defaults
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
