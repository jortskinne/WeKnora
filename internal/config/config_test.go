package config

import (
	"os"
	"testing"
)

func TestLoad_DefaultValues(t *testing.T) {
	os.Clearenv()

	cfg := Load()

	if cfg.AppName != "WeKnora" {
		t.Errorf("expected AppName 'WeKnora', got '%s'", cfg.AppName)
	}
	if cfg.AppEnv != "development" {
		t.Errorf("expected AppEnv 'development', got '%s'", cfg.AppEnv)
	}
	if cfg.Port != "8080" {
		t.Errorf("expected Port '8080', got '%s'", cfg.Port)
	}
	if cfg.Debug != false {
		t.Errorf("expected Debug false, got %v", cfg.Debug)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("expected LogLevel 'info', got '%s'", cfg.LogLevel)
	}
}

func TestLoad_EnvOverrides(t *testing.T) {
	t.Setenv("APP_NAME", "TestApp")
	t.Setenv("APP_ENV", "production")
	t.Setenv("PORT", "9090")
	t.Setenv("DEBUG", "true")
	t.Setenv("LOG_LEVEL", "debug")

	cfg := Load()

	if cfg.AppName != "TestApp" {
		t.Errorf("expected AppName 'TestApp', got '%s'", cfg.AppName)
	}
	if cfg.Port != "9090" {
		t.Errorf("expected Port '9090', got '%s'", cfg.Port)
	}
	if !cfg.Debug {
		t.Error("expected Debug to be true")
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("expected LogLevel 'debug', got '%s'", cfg.LogLevel)
	}
}

func TestConfig_IsDevelopment(t *testing.T) {
	cfg := &Config{AppEnv: "development"}
	if !cfg.IsDevelopment() {
		t.Error("expected IsDevelopment to return true")
	}
	if cfg.IsProduction() {
		t.Error("expected IsProduction to return false")
	}
}

func TestConfig_IsProduction(t *testing.T) {
	cfg := &Config{AppEnv: "production"}
	if !cfg.IsProduction() {
		t.Error("expected IsProduction to return true")
	}
	if cfg.IsDevelopment() {
		t.Error("expected IsDevelopment to return false")
	}
}

func TestLoad_InvalidDebugValue(t *testing.T) {
	t.Setenv("DEBUG", "notabool")

	cfg := Load()

	if cfg.Debug != false {
		t.Error("expected Debug to fall back to false on invalid value")
	}
}
