package sl

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

const (
	defaultBaseURL    = "https://secondlife.com"
	defaultFriendsURL = "/my/widget-friends.php"
	defaultGroupsURL  = "/my/widget-groups.php"
	defaultLindensURL = "/my/widget-linden-dollar.php"
	defaultTimeout    = 10 * time.Second
)

var defaultUserAgents = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0",
}

type Client struct {
	http       *http.Client
	token      string
	baseURL    string
	userAgents []string
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.http.Timeout = d }
}

func WithUserAgents(agents []string) Option {
	return func(c *Client) { c.userAgents = agents }
}

func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.http = hc }
}

func NewClient(token string, opts ...Option) *Client {
	c := &Client{
		http:       &http.Client{Timeout: defaultTimeout},
		token:      token,
		baseURL:    defaultBaseURL,
		userAgents: defaultUserAgents,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) fetch(ctx context.Context, path string) (io.ReadCloser, error) {
	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", c.randomUA())
	req.AddCookie(&http.Cookie{Name: "session-token", Value: c.token})
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request %s: %w", path, err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status %d for %s", resp.StatusCode, path)
	}
	return resp.Body, nil
}

func (c *Client) randomUA() string {
	return c.userAgents[rand.IntN(len(c.userAgents))]
}
