package tui

import (
	"fmt"
	"strings"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func renderGroupsPanel(groups []sl.Group, focused bool, width, height int) string {
	title := titleStyle.Render(fmt.Sprintf("Groups (%d)", len(groups)))

	var lines []string
	for _, g := range groups {
		lines = append(lines, fmt.Sprintf("  %s", g.Name))
	}

	if len(lines) == 0 {
		lines = append(lines, statusItemStyle.Render("  No groups to display"))
	}

	content := strings.Join(lines, "\n")

	style := panelStyle
	if focused {
		style = panelFocusedStyle
	}

	return style.Width(width).Height(height).Render(title + "\n" + content)
}
