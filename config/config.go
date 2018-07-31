package config

import "github.com/kelseyhightower/envconfig"

// StatsConfig ...
type StatsConfig struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

// GetConfig ...
func GetConfig() (*StatsConfig, error) {
	var c StatsConfig
	err := envconfig.Process("statsconfig", &c)
	return &c, err
}
