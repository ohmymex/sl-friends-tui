package config_test

import (
	"testing"
	"time"

	"github.com/ohmymex/sl-friends-tui/internal/config"
)

func TestLoad_Defaults(t *testing.T) {
	cfg := config.Default()
	if cfg.Filter != "online" {
		t.Errorf("expected filter 'online', got %q", cfg.Filter)
	}
	if cfg.Refresh != 5*time.Second {
		t.Errorf("expected refresh 5s, got %v", cfg.Refresh)
	}
	if cfg.ShowLindens != true {
		t.Error("expected show_lindens true by default")
	}
	if cfg.ShowGroups != true {
		t.Error("expected show_groups true by default")
	}
	if cfg.Layout != "dashboard" {
		t.Errorf("expected layout 'dashboard', got %q", cfg.Layout)
	}
	if cfg.Notify.Enabled != false {
		t.Error("expected notify disabled by default")
	}
}

func TestLoad_FromFile(t *testing.T) {
	cfg, err := config.LoadFromFile("testdata/valid.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Token != "abc-123-def" {
		t.Errorf("expected token 'abc-123-def', got %q", cfg.Token)
	}
	if cfg.TokenEncoded != true {
		t.Error("expected token_encoded true")
	}
	if cfg.Filter != "offline" {
		t.Errorf("expected filter 'offline', got %q", cfg.Filter)
	}
	if cfg.Refresh != 10*time.Second {
		t.Errorf("expected refresh 10s, got %v", cfg.Refresh)
	}
	if !cfg.Notify.Enabled {
		t.Error("expected notify enabled")
	}
	if len(cfg.Notify.Users) != 2 {
		t.Fatalf("expected 2 notify users, got %d", len(cfg.Notify.Users))
	}
}

func TestLoad_MinimalFile(t *testing.T) {
	cfg, err := config.LoadFromFile("testdata/minimal.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Token != "my-token" {
		t.Errorf("expected token 'my-token', got %q", cfg.Token)
	}
	if cfg.Filter != "online" {
		t.Errorf("expected default filter 'online', got %q", cfg.Filter)
	}
	if cfg.Refresh != 5*time.Second {
		t.Errorf("expected default refresh 5s, got %v", cfg.Refresh)
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := config.LoadFromFile("testdata/nonexistent.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
