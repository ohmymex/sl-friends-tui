package notify

import (
	"context"
	"fmt"
	"strings"
)

type MultiNotifier struct {
	notifiers []Notifier
}

func NewMultiNotifier(notifiers ...Notifier) *MultiNotifier {
	return &MultiNotifier{notifiers: notifiers}
}

func (m *MultiNotifier) Notify(ctx context.Context, title, body string) error {
	var errs []string
	for _, n := range m.notifiers {
		if err := n.Notify(ctx, title, body); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("notify errors: %s", strings.Join(errs, "; "))
	}
	return nil
}
