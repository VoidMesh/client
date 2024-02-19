package view

import (
	"context"
	"fmt"
	"strings"

	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/VoidMesh/client/internal/game"
	"github.com/VoidMesh/client/internal/program_context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	game game.Game
	ctx  *program_context.Ctx
}

func NewMainView(ctx *program_context.Ctx, g game.Game) MainView {
	v := MainView{ctx: ctx, game: g}

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

func (m MainView) View() string {
	s := strings.Builder{}
	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.sideBar(),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.NewStyle().Render("Welcome to Void Mesh!"),
				lipgloss.NewStyle().Render(fmt.Sprintf("You are playing as: %s", m.game.Character.Name)),
			),
		),
	)

	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (m MainView) sideBar() string {
	s := strings.Builder{}
	title := lipgloss.NewStyle().Underline(true).Bold(true).MarginBottom(1).Render(m.game.Character.Name)
	s.WriteString(title)
	s.WriteString("\n")

	return lipgloss.JoinVertical(
		lipgloss.Right,
		lipgloss.NewStyle().
			Width(((m.ctx.ScreenWidth)/24)*4).
			MaxWidth((((m.ctx.ScreenWidth)/24)*4)+2). // Margin + border
			Height(m.ctx.Height-3).                   // Newline - borders
			MaxHeight(m.ctx.Height+1).                // Bottom border
			Border(lipgloss.BlockBorder(), true, true, true, false).
			Margin(0, 1, 0, 0).
			Padding(1).
			Render(s.String()),
	)
}
