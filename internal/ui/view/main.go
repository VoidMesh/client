package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	model tea.Model
}

func NewMainView(m *tea.Model) MainView {
	v := MainView{model: *m}

	return v
}

func (v MainView) Init() tea.Cmd {
	return nil
}

func (v MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	return v, tea.Batch(cmds...)
}

func (v MainView) View() string {
	return lipgloss.NewStyle().Render("Main view")
}
