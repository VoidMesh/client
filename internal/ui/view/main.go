package view

import (
	"context"
	"fmt"
	"strings"

	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/VoidMesh/client/internal/game"
	"github.com/VoidMesh/client/internal/program_context"
	"github.com/VoidMesh/client/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	game    game.Game
	ctx     *program_context.Ctx
	currTab int
	tabs    []string
}

func NewMainView(ctx *program_context.Ctx, g game.Game) MainView {
	v := MainView{ctx: ctx, game: g}

	resp, _ := v.game.Services.Inventory.Read(context.TODO(), &inventory.ReadRequest{
		CharacterId: v.game.Character.Id,
	})

	v.game.Inventory = resp.Inventory

	// TODO: Create a view for each tab
	v.tabs = []string{
		"Skills",
		"Inventory",
		"Ships",
		"Quests",
		"Map",
		"Industrial",
		"Credits",
	}

	return v
}

func (v MainView) Init() tea.Cmd {
	return nil
}

func (v MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.PrevTab):
			v.setCurrentTab(v.getPrevTabId())

		case key.Matches(msg, utils.Keys.NextTab):
			v.setCurrentTab(v.getNextTabId())
		}
	}

	cmds = append(cmds, cmd)
	return v, tea.Batch(cmds...)
}

func (v MainView) View() string {
	s := strings.Builder{}
	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		v.sideBar(),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.NewStyle().Render("Welcome to Void Mesh!"),
				lipgloss.NewStyle().Render(fmt.Sprintf("You are playing as: %s", v.game.Character.Name)),
			),
		),
	)

	s.WriteString(mainContent)
	s.WriteString("\n")
	return s.String()
}

func (v MainView) sideBar() string {
	s := strings.Builder{}
	title := lipgloss.NewStyle().Underline(true).Bold(true).MarginBottom(1).Render(v.game.Character.Name)
	s.WriteString(title)
	s.WriteString("\n")

	return lipgloss.JoinVertical(
		lipgloss.Right,
		lipgloss.NewStyle().
			Width(((v.ctx.ScreenWidth)/24)*4).
			MaxWidth((((v.ctx.ScreenWidth)/24)*4)+2). // Margin + border
			Height(v.ctx.Height-3).                   // Newline - borders
			MaxHeight(v.ctx.Height+1).                // Bottom border
			Border(lipgloss.BlockBorder(), true, true, true, false).
			Margin(0, 1, 0, 0).
			Padding(1).
			Render(
				s.String(),
				v.TabsList(),
			),
	)
}

func (v MainView) TabsList() string {
	var tabsTitle []string
	var titleStyle lipgloss.Style

	for i, tab := range v.tabs {
		titleStyle = lipgloss.NewStyle()
		if v.currTab == i {
			titleStyle = titleStyle.Copy().Bold(true).Underline(true)
		}
		tabsTitle = append(tabsTitle, titleStyle.Render(tab))
	}
	return lipgloss.JoinVertical(lipgloss.Left, tabsTitle...)
}

func (v *MainView) setCurrentTab(tab int) {
	v.currTab = tab
}

func (v MainView) getTabAt(id int) string {
	tabs := v.tabs
	return tabs[id]
}

func (v MainView) getPrevTabId() int {
	var currTabId int
	tabs := v.tabs
	currTabId = (v.currTab - 1) % len(tabs)
	if currTabId < 0 {
		currTabId += len(tabs)
	}

	return currTabId
}

func (v MainView) getNextTabId() int {
	return (v.currTab + 1) % len(v.tabs)
}
