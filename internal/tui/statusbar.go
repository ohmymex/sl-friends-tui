package tui

import (
	"fmt"
	"strings"
	"time"
)

func renderStatusBar(lindens string, filter string, refresh time.Duration, err error, searching bool, width int) string {
	var parts []string

	if lindens != "" {
		parts = append(parts, lindens)
	}

	parts = append(parts, fmt.Sprintf("%s %s",
		statusItemStyle.Render("Filter:"),
		statusValueStyle.Render(filter),
	))

	parts = append(parts, fmt.Sprintf("%s %s",
		statusItemStyle.Render("Refresh:"),
		statusValueStyle.Render(refresh.String()),
	))

	if err != nil {
		parts = append(parts, errorStyle.Render(fmt.Sprintf("Error: %s", err)))
	}

	if searching {
		parts = append(parts, helpKeyStyle.Render("ESC")+helpDescStyle.Render(" cancel"))
	} else {
		parts = append(parts, helpKeyStyle.Render("?")+helpDescStyle.Render(" help"))
	}

	content := strings.Join(parts, statusItemStyle.Render("  │  "))

	return statusBarStyle.Width(width).Render(content)
}

func renderHelp() string {
	keys := []struct{ key, desc string }{
		{"Tab", "Switch pane"},
		{"j/k", "Scroll down/up"},
		{"/", "Search friends"},
		{"Esc", "Cancel search"},
		{"f", "Cycle filter"},
		{"r", "Refresh now"},
		{"?", "Toggle help"},
		{"q", "Quit"},
	}

	var lines []string
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("  %s  %s",
			helpKeyStyle.Render(fmt.Sprintf("%-5s", k.key)),
			helpDescStyle.Render(k.desc),
		))
	}

	return panelStyle.Render(
		titleStyle.Render("Keyboard Shortcuts") + "\n\n" + strings.Join(lines, "\n"),
	)
}
