package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func renderFriendsPanel(friends []sl.Friend, filter string, searchQuery string, focused bool, showInternal bool, width, height int, scroll int) string {
	filtered := filterFriends(friends, filter, searchQuery)

	onlineCount := 0
	for _, f := range filtered {
		if f.Online {
			onlineCount++
		}
	}

	title := titleStyle.Render(fmt.Sprintf("Friends Online (%d)", onlineCount))
	if filter == "all" {
		title = titleStyle.Render(fmt.Sprintf("Friends (%d/%d online)", onlineCount, len(filtered)))
	} else if filter == "offline" {
		title = titleStyle.Render(fmt.Sprintf("Friends Offline (%d)", len(filtered)))
	}

	var lines []string
	for _, f := range filtered {
		lines = append(lines, formatFriend(f, showInternal))
	}

	if len(lines) == 0 {
		lines = append(lines, statusItemStyle.Render("  No friends to display"))
	}

	if scroll > len(lines) {
		scroll = len(lines)
	}
	visible := lines[scroll:]

	maxVisible := height - 3
	if maxVisible < 1 {
		maxVisible = 1
	}
	if len(visible) > maxVisible {
		visible = visible[:maxVisible]
	}

	content := strings.Join(visible, "\n")

	style := panelStyle
	if focused {
		style = panelFocusedStyle
	}

	return style.Width(width).Height(height).Render(title + "\n" + content)
}

func formatFriend(f sl.Friend, showInternal bool) string {
	var dot string
	if f.Online {
		dot = onlineStyle.Render(onlineDot)
	} else {
		dot = offlineStyle.Render(offlineDot)
	}

	name := f.DisplayName
	if showInternal && f.InternalName != "" {
		name = fmt.Sprintf("%s (%s)", f.DisplayName, f.InternalName)
	}

	if f.Online {
		return fmt.Sprintf("  %s%s", dot, name)
	}
	return fmt.Sprintf("  %s%s", dot, lipgloss.NewStyle().Foreground(colorOffline).Render(name))
}

func filterFriends(friends []sl.Friend, filter string, searchQuery string) []sl.Friend {
	var result []sl.Friend
	query := strings.ToLower(searchQuery)

	for _, f := range friends {
		switch filter {
		case "online":
			if !f.Online {
				continue
			}
		case "offline":
			if f.Online {
				continue
			}
		}

		if query != "" {
			nameMatch := strings.Contains(strings.ToLower(f.DisplayName), query)
			internalMatch := strings.Contains(strings.ToLower(f.InternalName), query)
			if !nameMatch && !internalMatch {
				continue
			}
		}

		result = append(result, f)
	}

	return result
}
