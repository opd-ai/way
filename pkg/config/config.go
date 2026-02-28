package config

import (
	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	Window  WindowConfig  `mapstructure:"window"`
	Server  ServerConfig  `mapstructure:"server"`
	Game    GameConfig    `mapstructure:"game"`
}

// WindowConfig holds display settings.
type WindowConfig struct {
	Width  int    `mapstructure:"width"`
	Height int    `mapstructure:"height"`
	Title  string `mapstructure:"title"`
}

// ServerConfig holds network server settings.
type ServerConfig struct {
	Address    string `mapstructure:"address"`
	TickRateHz int    `mapstructure:"tick_rate_hz"`
}

// GameConfig holds gameplay settings.
type GameConfig struct {
	Seed       int64   `mapstructure:"seed"`
	Genre      int     `mapstructure:"genre"`
	Difficulty float64 `mapstructure:"difficulty"`
}

// Load reads configuration from file and environment, applying defaults.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.way")

	viper.SetDefault("window.width", 800)
	viper.SetDefault("window.height", 600)
	viper.SetDefault("window.title", "Way")
	viper.SetDefault("server.address", ":7777")
	viper.SetDefault("server.tick_rate_hz", 20)
	viper.SetDefault("game.seed", 0)
	viper.SetDefault("game.genre", 0)
	viper.SetDefault("game.difficulty", 1.0)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
