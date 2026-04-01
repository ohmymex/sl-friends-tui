package tui

import (
	"context"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohmymex/sl-friends-tui/internal/config"
	"github.com/ohmymex/sl-friends-tui/internal/notify"
	"github.com/ohmymex/sl-friends-tui/pkg/sl"
)

type Pane int

const (
	PaneFriends Pane = iota + 1
	PaneGroups
)

type (
	TickMsg          time.Time
	FriendsResultMsg struct {
		friends []sl.Friend
		err     error
	}
	GroupsResultMsg struct {
		groups []sl.Group
		err    error
	}
	LindensResultMsg struct {
		lindens *sl.Lindens
		err     error
	}
)

type App struct {
	client   *sl.Client
	config   *config.Config
	notifier notify.Notifier

	friends []sl.Friend
	groups  []sl.Group
	lindens *sl.Lindens
	err     error

	activePane    Pane
	search        textinput.Model
	searching     bool
	showHelp      bool
	width         int
	height        int
	friendsScroll int
	groupsScroll  int

	notified   map[string]bool
	prevOnline map[string]bool
}

func New(client *sl.Client, cfg *config.Config, notifier notify.Notifier) *App {
	return &App{
		client:     client,
		config:     cfg,
		notifier:   notifier,
		activePane: PaneFriends,
		search:     newSearchInput(),
		notified:   make(map[string]bool),
		prevOnline: make(map[string]bool),
	}
}

func (a *App) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(a.config.Refresh),
		fetchAllCmd(a.client),
	)
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return a.handleKey(msg)

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		return a, nil

	case TickMsg:
		return a, tea.Batch(
			tickCmd(a.config.Refresh),
			fetchAllCmd(a.client),
		)

	case FriendsResultMsg:
		if msg.err != nil {
			a.err = msg.err
		} else {
			a.err = nil
			a.checkNotifications(msg.friends)
			a.friends = msg.friends
		}
		return a, nil

	case GroupsResultMsg:
		if msg.err != nil {
			a.err = msg.err
		} else {
			a.groups = msg.groups
		}
		return a, nil

	case LindensResultMsg:
		if msg.err != nil {
			a.err = msg.err
		} else {
			a.lindens = msg.lindens
		}
		return a, nil
	}

	if a.searching {
		var cmd tea.Cmd
		a.search, cmd = a.search.Update(msg)
		return a, cmd
	}

	return a, nil
}

func (a *App) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return a, tea.Quit
	}

	if a.searching {
		switch msg.String() {
		case "esc":
			a.searching = false
			a.search.SetValue("")
			a.search.Blur()
			return a, nil
		case "enter":
			a.searching = false
			a.search.Blur()
			return a, nil
		}
		var cmd tea.Cmd
		a.search, cmd = a.search.Update(msg)
		return a, cmd
	}

	switch msg.String() {
	case "q":
		return a, tea.Quit
	case "tab":
		if a.activePane == PaneFriends {
			a.activePane = PaneGroups
		} else {
			a.activePane = PaneFriends
		}
		return a, nil
	case "/":
		a.searching = true
		a.search.Focus()
		return a, textinput.Blink
	case "f":
		a.cycleFilter()
		return a, nil
	case "r":
		return a, fetchAllCmd(a.client)
	case "?":
		a.showHelp = !a.showHelp
		return a, nil
	case "j", "down":
		a.scrollDown()
		return a, nil
	case "k", "up":
		a.scrollUp()
		return a, nil
	}

	return a, nil
}

func (a *App) scrollDown() {
	if a.activePane == PaneFriends {
		maxScroll := len(filterFriends(a.friends, a.config.Filter, a.search.Value())) - a.panelContentHeight()
		if maxScroll < 0 {
			maxScroll = 0
		}
		if a.friendsScroll < maxScroll {
			a.friendsScroll++
		}
	} else {
		maxScroll := len(a.groups) - a.panelContentHeight()
		if maxScroll < 0 {
			maxScroll = 0
		}
		if a.groupsScroll < maxScroll {
			a.groupsScroll++
		}
	}
}

func (a *App) scrollUp() {
	if a.activePane == PaneFriends {
		if a.friendsScroll > 0 {
			a.friendsScroll--
		}
	} else {
		if a.groupsScroll > 0 {
			a.groupsScroll--
		}
	}
}

func (a *App) panelContentHeight() int {
	statusBarHeight := 3
	searchBarHeight := 0
	if a.searching {
		searchBarHeight = 3
	}
	h := a.height - statusBarHeight - searchBarHeight - 5
	if h < 1 {
		h = 1
	}
	return h
}

func (a *App) cycleFilter() {
	a.friendsScroll = 0
	switch a.config.Filter {
	case "online":
		a.config.Filter = "offline"
	case "offline":
		a.config.Filter = "all"
	case "all":
		a.config.Filter = "online"
	}
}

func (a *App) checkNotifications(newFriends []sl.Friend) {
	if a.notifier == nil || !a.config.Notify.Enabled || len(a.config.Notify.Users) == 0 {
		return
	}

	watched := make(map[string]bool)
	for _, u := range a.config.Notify.Users {
		watched[u] = true
	}

	for _, f := range newFriends {
		if !watched[f.InternalName] {
			continue
		}

		wasOnline := a.prevOnline[f.InternalName]

		if f.Online && !wasOnline && !a.notified[f.InternalName] {
			_ = a.notifier.Notify(context.Background(),
				"Friend Online",
				f.DisplayName+" just connected!")
			a.notified[f.InternalName] = true
		}

		if !f.Online {
			a.notified[f.InternalName] = false
		}
	}

	a.prevOnline = make(map[string]bool)
	for _, f := range newFriends {
		a.prevOnline[f.InternalName] = f.Online
	}
}

func (a *App) View() string {
	if a.width == 0 || a.height == 0 {
		return "Loading..."
	}

	if a.showHelp {
		return lipgloss.Place(a.width, a.height,
			lipgloss.Center, lipgloss.Center,
			renderHelp(),
		)
	}

	statusBarHeight := 3
	searchBarHeight := 0
	if a.searching {
		searchBarHeight = 3
	}
	panelHeight := a.height - statusBarHeight - searchBarHeight - 2
	if panelHeight < 3 {
		panelHeight = 3
	}

	friendsPanelWidth := a.width*2/3 - 2
	groupsPanelWidth := a.width - friendsPanelWidth - 4

	friendsPanel := renderFriendsPanel(
		a.friends,
		a.config.Filter,
		a.search.Value(),
		a.activePane == PaneFriends,
		a.config.ShowInternal,
		friendsPanelWidth,
		panelHeight,
		a.friendsScroll,
	)

	groupsPanel := renderGroupsPanel(
		a.groups,
		a.activePane == PaneGroups,
		groupsPanelWidth,
		panelHeight,
		a.groupsScroll,
	)

	panels := lipgloss.JoinHorizontal(lipgloss.Top, friendsPanel, groupsPanel)

	lindensStr := ""
	if a.config.ShowLindens {
		lindensStr = renderLindens(a.lindens)
	}
	statusBar := renderStatusBar(lindensStr, a.config.Filter, a.config.Refresh, a.err, a.searching, a.width-4)

	var sections []string
	sections = append(sections, panels)
	if a.searching {
		sections = append(sections, renderSearch(a.search, a.width-4))
	}
	sections = append(sections, statusBar)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func fetchAllCmd(client *sl.Client) tea.Cmd {
	return tea.Batch(
		fetchFriendsCmd(client),
		fetchGroupsCmd(client),
		fetchLindensCmd(client),
	)
}

func fetchFriendsCmd(client *sl.Client) tea.Cmd {
	return func() tea.Msg {
		friends, err := client.FetchFriends(context.Background())
		return FriendsResultMsg{friends: friends, err: err}
	}
}

func fetchGroupsCmd(client *sl.Client) tea.Cmd {
	return func() tea.Msg {
		groups, err := client.FetchGroups(context.Background())
		return GroupsResultMsg{groups: groups, err: err}
	}
}

func fetchLindensCmd(client *sl.Client) tea.Cmd {
	return func() tea.Msg {
		lindens, err := client.FetchLindens(context.Background())
		return LindensResultMsg{lindens: lindens, err: err}
	}
}
