package sl

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) FetchLindens(ctx context.Context) (*Lindens, error) {
	body, err := c.fetch(ctx, defaultLindensURL)
	if err != nil {
		return nil, fmt.Errorf("fetch lindens: %w", err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("parse lindens html: %w", err)
	}

	balance := ""
	doc.Find(".main-widget-content strong").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			balance = text
		}
	})

	return &Lindens{Balance: balance}, nil
}
