package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	cfg := Load()
	if cfg.Port == "" {
		t.Fatal("port should default")
	}
	if cfg.ReadTimeout <= 0 {
		t.Fatal("read timeout should default")
	}
}