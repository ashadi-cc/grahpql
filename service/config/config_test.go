package config

import (
	"os"
	"testing"
)

func TestEnvFromProcess(t *testing.T) {
	_ = os.Setenv("DEBUG_MODE", "false")
	config := GetConfig()
	if config.DebugMode {
		t.Fatalf("debug mode should false but got %+v", config)
	}
}

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	if config.AppName != "grahpql" {
		t.Fatalf("read config file vailed %+v", config)
	}
}
