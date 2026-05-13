package views

import tea "github.com/charmbracelet/bubbletea"

type ListRoomView struct{}

func NewListRoomView() *ListRoomView {
	return &ListRoomView{}
}

func (lrv *ListRoomView) SetQuery(q string) {

}

func (lrv *ListRoomView) Update(tea.Msg) (View, tea.Cmd) {
	return lrv, nil
}

func (lrv *ListRoomView) MainView() string {
	return ""
}

func (lrv *ListRoomView) MenuView() string {
	return ""
}
