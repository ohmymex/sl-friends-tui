package notify

import (
	"context"

	"github.com/gen2brain/beeep"
)

// DesktopNotifier sends cross-platform desktop notifications.
type DesktopNotifier struct{}

func NewDesktopNotifier() *DesktopNotifier {
	return &DesktopNotifier{}
}

func (d *DesktopNotifier) Notify(_ context.Context, title, body string) error {
	return beeep.Notify(title, body, "")
}
