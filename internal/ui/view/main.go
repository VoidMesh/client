package view

import (
	"context"
	"fmt"

	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/VoidMesh/client/internal/game"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	game game.Game
}

func NewMainView(g game.Game) MainView {
	v := MainView{game: g}

	resp, _ := v.game.Services.Inventory.Read(context.TODO(), &inventory.ReadRequest{
		CharacterId: v.game.Character.Id,
	})

	g.Inventory = resp.Inventory

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
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.NewStyle().Render("Welcome to Void Mesh!"),
			lipgloss.NewStyle().Render(fmt.Sprintf("You are playing as: %s", v.game.Character.Name)),
		),
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.NewStyle().Render("Inventory"),
			lipgloss.NewStyle().Render(fmt.Sprintf("Inventory %v:", v.game.Inventory)),
		),
	)
}
