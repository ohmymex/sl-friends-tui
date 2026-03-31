package sl_test

import (
	"testing"
	"time"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func TestNewClient_Defaults(t *testing.T) {
	c := sl.NewClient("test-token-123")
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNewClient_WithTimeout(t *testing.T) {
	c := sl.NewClient("test-token-123", sl.WithTimeout(30*time.Second))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNewClient_WithUserAgents(t *testing.T) {
	agents := []string{"Agent/1.0", "Agent/2.0"}
	c := sl.NewClient("test-token-123", sl.WithUserAgents(agents))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}
