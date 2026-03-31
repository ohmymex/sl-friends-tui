package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

func newSearchInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Search friends..."
	ti.Prompt = searchPromptStyle.Render("/ ")
	ti.CharLimit = 64
	return ti
}

func renderSearch(input textinput.Model, width int) string {
	return panelStyle.Width(width).Render(fmt.Sprintf("%s", input.View()))
}
