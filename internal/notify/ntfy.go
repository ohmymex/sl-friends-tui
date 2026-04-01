package notify

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type NtfyNotifier struct {
	server   string
	topic    string
	token    string
	priority string
	client   *http.Client
}

type NtfyOption func(*NtfyNotifier)

func WithNtfyToken(token string) NtfyOption {
	return func(n *NtfyNotifier) { n.token = token }
}

func WithNtfyPriority(p int) NtfyOption {
	return func(n *NtfyNotifier) { n.priority = fmt.Sprintf("%d", p) }
}

func NewNtfyNotifier(server, topic string, opts ...NtfyOption) *NtfyNotifier {
	n := &NtfyNotifier{
		server:   strings.TrimRight(server, "/"),
		topic:    topic,
		priority: "3",
		client:   &http.Client{},
	}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func (n *NtfyNotifier) Notify(ctx context.Context, title, body string) error {
	url := fmt.Sprintf("%s/%s", n.server, n.topic)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("ntfy request: %w", err)
	}

	req.Header.Set("Title", title)
	req.Header.Set("Priority", n.priority)
	req.Header.Set("Tags", "green_circle")

	if n.token != "" {
		req.Header.Set("Authorization", "Bearer "+n.token)
	}

	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("ntfy send: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ntfy unexpected status: %d", resp.StatusCode)
	}

	return nil
}
