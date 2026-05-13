package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fajarilf/grpc-chat-client/app/utils"
)

type styles struct {
	title        lipgloss.Style
	item         lipgloss.Style
	selectedItem lipgloss.Style
	muted        lipgloss.Style
}

func NewStyles(darkBG bool) styles {
	return styles{
		title:        lipgloss.NewStyle().MarginLeft(2).Bold(true),
		item:         lipgloss.NewStyle().PaddingLeft(4),
		selectedItem: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")),
		muted:        lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

var menuKeys = struct {
	Up, Down, Enter key.Binding
}{
	Up:    key.NewBinding(key.WithKeys("up", "k")),
	Down:  key.NewBinding(key.WithKeys("down", "j")),
	Enter: key.NewBinding(key.WithKeys("enter")),
}

var itemList = []string{
	"create room", "join room", "list room",
}

type MenuView struct {
	itemList     []string
	selected     int
	open         bool
	lastSelected string
	styles       styles
}

func NewMenuView() *MenuView {
	return &MenuView{
		itemList: itemList,
		selected: 0,
		open:     false,
		styles:   NewStyles(true),
	}
}

// SetQuery is called by the parent to derive menu visibility from the
// textinput's current value. The menu is open iff the value is exactly "/".
func (mv *MenuView) SetQuery(q string) {
	wasOpen := mv.open
	if len(q) == 0 {
		mv.open = false
		return
	}

	beginWith := q[0]
	mv.open = beginWith == '/'

	mv.itemList = utils.SearchQuery(itemList, q)

	if mv.open && !wasOpen {
		mv.selected = 0
	}
}

func (mv *MenuView) Update(msg tea.Msg) (View, tea.Cmd) {
	km, ok := msg.(tea.KeyMsg)
	if !ok || !mv.open {
		return mv, nil
	}

	switch {
	case key.Matches(km, menuKeys.Up):
		if mv.selected > 0 {
			mv.selected--
		}
	case key.Matches(km, menuKeys.Down):
		if mv.selected < len(mv.itemList)-1 {
			mv.selected++
		}
	case key.Matches(km, menuKeys.Enter):
		// selected := mv.itemList[mv.selected]
		mv.open = false
		// return mv, Navigate(selected)
	}
	return mv, nil
}

func (mv *MenuView) MainView() string {
	var b strings.Builder

	if !mv.open {
		if mv.lastSelected != "" {
			b.WriteString("Selected: " + mv.lastSelected + "\n")
			b.WriteString(mv.styles.muted.Render("  ⎿  done") + "\n")
		} else {
			b.WriteString("\n")
		}
		return b.String()
	}

	return b.String()
}

func (mv *MenuView) MenuView() string {
	var b strings.Builder

	if !mv.open {
		return b.String()
	}

	for i, item := range mv.itemList {
		if i == mv.selected {
			b.WriteString(mv.styles.selectedItem.Render("> " + item))
		} else {
			b.WriteString(mv.styles.item.Render(item))
		}
		b.WriteString("\n")
	}
	return b.String()
}
