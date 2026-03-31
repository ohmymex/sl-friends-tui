package notify

import "context"

// Notifier sends notifications. Implementations may use desktop notifications,
// push services, or other backends.
type Notifier interface {
	Notify(ctx context.Context, title string, body string) error
}
