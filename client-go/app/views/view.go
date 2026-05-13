package views

import tea "github.com/charmbracelet/bubbletea"

type View interface {
	SetQuery(string)
	Update(tea.Msg) (View, tea.Cmd)
	MainView() string
	MenuView() string
}
