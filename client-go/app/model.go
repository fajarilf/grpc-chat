package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fajarilf/grpc-chat-client/app/views"
)

var inputBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderLeft(false).BorderRight(false).
	BorderForeground(lipgloss.Color("63")).
	Padding(0, 1)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	Exit key.Binding
	Help key.Binding
}

// ShortHelp is the one-line hint rendered by help.Model.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Exit, k.Help}
}

// FullHelp is the expanded multi-column view (shown when help is toggled).
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Exit, k.Help},
	}
}

var defaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Exit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit app"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}

type Model struct {
	textInput textinput.Model
	help      help.Model
	current   views.View
	width     int
}

func NewModel() *Model {
	ti := textinput.New()
	ti.Placeholder = "/ to open menu"
	ti.Focus()

	return &Model{
		textInput: ti,
		help:      help.New(),
		current:   views.NewMenuView(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.current.SetQuery(m.textInput.Value())

	switch msg := msg.(type) {
	case views.NavigateMsg:
		switch msg.Dest {
		case "create room":
		case "join room":
		case "list room":
			m.current = views.NewListRoomView()
		}
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.width = msg.Width
		m.textInput.Width = msg.Width - 4
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultKeyMap.Exit):
			return m, tea.Quit
		case key.Matches(msg, defaultKeyMap.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}
	}

	next, cmd := m.current.Update(msg)
	cmds = append(cmds, cmd)
	m.current = next

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	box := inputBoxStyle.Render(m.textInput.View())
	footer := ""

	if len(m.current.MenuView()) != 0 {
		footer = m.current.MenuView()
	} else {
		footer = m.help.View(defaultKeyMap)
	}

	return m.current.MainView() + "\n" + box + "\n" + footer
}
