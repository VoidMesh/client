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

	v.game.Inventory = resp.Inventory

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
	var slotsString []string
	for _, s := range v.game.Inventory.Slots {
		slotsString = append(slotsString, fmt.Sprintf("%s: %d", s.Resource.Name, s.Quantity))
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Render("Welcome to Void Mesh!"),
			lipgloss.NewStyle().Render(fmt.Sprintf("You are playing as: %s", v.game.Character.Name)),
		),
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Render("Inventory"),
			lipgloss.JoinVertical(
				lipgloss.Left,
				slotsString...,
			),
		),
	)
}
