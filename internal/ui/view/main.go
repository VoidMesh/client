package view

import (
	"fmt"

	"github.com/VoidMesh/client/internal/game"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	game game.Game
}

func NewMainView(g game.Game) MainView {
	v := MainView{game: g}

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
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().Render("Welcome to Void Mesh!"),
		lipgloss.NewStyle().Render(fmt.Sprintf("You are playing as: %s", v.game.Character.Name)),
	)
}
