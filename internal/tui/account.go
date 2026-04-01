package tui

import (
	"fmt"
	"strings"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func renderAccountBar(account *sl.Account, width int) string {
	if account == nil {
		return ""
	}

	var parts []string

	if account.Username != "" {
		parts = append(parts, fmt.Sprintf("%s %s",
			titleStyle.Render(account.Username),
			statusItemStyle.Render(fmt.Sprintf("(%s)", account.Plan)),
		))
	}

	if account.Status != "" {
		parts = append(parts, fmt.Sprintf("%s %s",
			statusItemStyle.Render("Status:"),
			statusValueStyle.Render(account.Status),
		))
	}

	if account.LBalance != "" {
		parts = append(parts, statusValueStyle.Render(account.LBalance))
	}

	if account.USDBalance != "" {
		parts = append(parts, statusValueStyle.Render(account.USDBalance))
	}

	if account.Country != "" {
		parts = append(parts, fmt.Sprintf("%s %s",
			statusItemStyle.Render("Country:"),
			statusValueStyle.Render(account.Country),
		))
	}

	content := strings.Join(parts, statusItemStyle.Render("  │  "))
	return statusBarStyle.Width(width).Render(content)
}
