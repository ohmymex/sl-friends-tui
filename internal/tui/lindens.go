package tui

import (
	"fmt"

	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

func renderLindens(lindens *sl.Lindens) string {
	if lindens == nil || lindens.Balance == "" {
		return ""
	}
	return fmt.Sprintf("%s %s",
		statusItemStyle.Render("L$"),
		statusValueStyle.Render(lindens.Balance),
	)
}
