package notify_test

import (
	"context"
	"testing"

	"github.com/ohmymex/sl-friends-tui/internal/notify"
)

type mockNotifier struct {
	calls []struct{ title, body string }
}

func (m *mockNotifier) Notify(_ context.Context, title, body string) error {
	m.calls = append(m.calls, struct{ title, body string }{title, body})
	return nil
}

func TestMockNotifier_Implements_Interface(t *testing.T) {
	var n notify.Notifier = &mockNotifier{}
	err := n.Notify(context.Background(), "title", "body")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDesktopNotifier_Implements_Interface(t *testing.T) {
	var _ notify.Notifier = notify.NewDesktopNotifier()
}
