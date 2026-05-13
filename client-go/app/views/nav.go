package views

import tea "github.com/charmbracelet/bubbletea"

type NavigateMsg struct{ Dest string }

func Navigate(Dest string) tea.Cmd {
	return func() tea.Msg { return NavigateMsg{Dest} }
}
