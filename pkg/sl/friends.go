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
	doc.Find("#widgetFriendsOnlineContent .trigger").Each(func(i int, s *goquery.Selection) {
		span := s.Find("span")
		internalName := strings.TrimSuffix(span.AttrOr("title", ""), " Resident")
		friend := Friend{
			DisplayName:  strings.TrimSpace(span.Text()),
			InternalName: internalName,
			Online:       s.HasClass("online"),
		}
		friends = append(friends, friend)
	})

	return friends, nil
}
