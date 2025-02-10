// Package config loads runtime configuration from the environment.
package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all tunable server settings.
type Config struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// Load reads configuration from environment variables, applying sane defaults.
func Load() Config {
	return Config{
		Port:            getenv("PORT", "8080"),
		ReadTimeout:     getdur("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getdur("WRITE_TIMEOUT", 10*time.Second),
		ShutdownTimeout: getdur("SHUTDOWN_TIMEOUT", 15*time.Second),
	}
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getdur(key string, fallback time.Duration) time.Duration {
	if v, ok := os.LookupEnv(key); ok {
		if secs, err := strconv.Atoi(v); err == nil {
			return time.Duration(secs) * time.Second
		}
	}
	return fallback
}
