package sl

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) FetchFriends(ctx context.Context) ([]Friend, error) {
	body, err := c.fetch(ctx, defaultFriendsURL)
	if err != nil {
		return nil, fmt.Errorf("fetch friends: %w", err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("parse friends html: %w", err)
	}

	var friends []Friend
	doc.Find("#widgetFriendsOnlineContent tr.friend-status").Each(func(i int, s *goquery.Selection) {
		nameTd := s.Find("td.friend")
		statusTd := s.Find("td").Not(".friend")

		nameSpan := nameTd.Find("span").First()
		displayName := strings.TrimSpace(nameSpan.Text())
		internalName := strings.TrimSuffix(nameSpan.AttrOr("title", ""), " Resident")
		online := statusTd.HasClass("online")

		if displayName != "" {
			friends = append(friends, Friend{
				DisplayName:  displayName,
				InternalName: internalName,
				Online:       online,
			})
		}
	})

	return friends, nil
}
