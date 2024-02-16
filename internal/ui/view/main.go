package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	View *Model
}

func NewMainView() MainView {
	m := MainView{}

	return m
}

func (v MainView) Init() tea.Cmd {
	return nil
}

func (v MainView) Update(msg tea.Msg) (View, tea.Cmd) {
	var cmds []tea.Cmd

	return v, tea.Batch(cmds...)
}

func (v MainView) Render() string {
	return lipgloss.NewStyle().Render("Main view")
}
