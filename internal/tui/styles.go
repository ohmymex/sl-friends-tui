package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorPrimary   = lipgloss.Color("#7B61FF")
	colorSecondary = lipgloss.Color("#36CFC9")
	colorOnline    = lipgloss.Color("#52C41A")
	colorOffline   = lipgloss.Color("#8C8C8C")
	colorError     = lipgloss.Color("#FF4D4F")
	colorBorder    = lipgloss.Color("#434343")
	colorTitle     = lipgloss.Color("#FFFFFF")
	colorMuted     = lipgloss.Color("#8C8C8C")

	panelStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorBorder).
		Padding(0, 1)

	panelFocusedStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary).
		Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(colorTitle)

	onlineStyle = lipgloss.NewStyle().
		Foreground(colorOnline)

	offlineStyle = lipgloss.NewStyle().
		Foreground(colorOffline)

	statusBarStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorBorder).
		Padding(0, 1)

	statusItemStyle = lipgloss.NewStyle().
		Foreground(colorMuted)

	statusValueStyle = lipgloss.NewStyle().
		Foreground(colorSecondary)

	searchPromptStyle = lipgloss.NewStyle().
		Foreground(colorPrimary)

	errorStyle = lipgloss.NewStyle().
		Foreground(colorError)

	helpKeyStyle = lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true)

	helpDescStyle = lipgloss.NewStyle().
		Foreground(colorMuted)
)

const (
	onlineDot  = "● "
	offlineDot = "○ "
)
