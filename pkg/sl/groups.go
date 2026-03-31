package sl

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) FetchGroups(ctx context.Context) ([]Group, error) {
	body, err := c.fetch(ctx, defaultGroupsURL)
	if err != nil {
		return nil, fmt.Errorf("fetch groups: %w", err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("parse groups html: %w", err)
	}

	var groups []Group
	doc.Find("#widgetGroupsContent .trigger").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find("span").Text())
		if name != "" {
			groups = append(groups, Group{Name: name})
		}
	})

	return groups, nil
}
