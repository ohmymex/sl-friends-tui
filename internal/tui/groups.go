package tui

import (
	"fmt"
	"strings"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func renderGroupsPanel(groups []sl.Group, focused bool, width, height int, scroll int) string {
	title := titleStyle.Render(fmt.Sprintf("Groups (%d)", len(groups)))

	var lines []string
	for _, g := range groups {
		if g.MemberCount != "" {
			lines = append(lines, fmt.Sprintf("  %s (%s)", g.Name, g.MemberCount))
		} else {
			lines = append(lines, fmt.Sprintf("  %s", g.Name))
		}
	}

	if len(lines) == 0 {
		lines = append(lines, statusItemStyle.Render("  No groups to display"))
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
