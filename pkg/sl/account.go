package sl

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *Client) FetchAccount(ctx context.Context) (*Account, error) {
	body, err := c.fetch(ctx, defaultAccountURL)
	if err != nil {
		return nil, fmt.Errorf("fetch account: %w", err)
	}
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("parse account html: %w", err)
	}

	account := &Account{}
	account.Username = strings.TrimSpace(doc.Find("#resident-name").Text())

	doc.Find(".account-summary-table tr").Each(func(i int, s *goquery.Selection) {
		th := strings.TrimSpace(s.Find("th").Text())
		td := s.Find("td")

		switch {
		case strings.Contains(th, "current plan"):
			account.Plan = strings.TrimSpace(td.Find("strong").Text())
		case strings.Contains(th, "current status"):
			account.Status = cleanText(td.Text())
		case strings.Contains(th, "Country"):
			account.Country = cleanText(strings.Split(td.Text(), " Change")[0])
		case strings.Contains(th, "L$ Balance"):
			account.LBalance = cleanBalance(td.Text())
		case strings.Contains(th, "US Dollar Balance"):
			account.USDBalance = strings.TrimSpace(td.Find(".usd-balance-text").Text())
		case strings.Contains(th, "Current holdings"):
			account.LandCurrent = cleanText(td.Text())
		}
	})

	return account, nil
}

func cleanText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func cleanBalance(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\u200a", " ")
	s = strings.ReplaceAll(s, "&hairsp;", " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}
