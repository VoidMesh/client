package view

import (
	"strings"

	"github.com/VoidMesh/client/internal/game"
	"github.com/VoidMesh/client/internal/ui"
	"github.com/VoidMesh/client/internal/ui/tab"
	"github.com/VoidMesh/client/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	game      game.Game
	ctx       *ui.Context
	currTabId int
	currTab   tab.Tab
	tabs      []tab.Tab
}

func NewMainView(ctx *ui.Context, g game.Game) MainView {
	v := MainView{ctx: ctx, game: g}

	v.tabs = []tab.Tab{
		tab.NewSkillsTab(),
		tab.NewInventoryTab(),
	}
	v.currTab = v.tabs[0]

	for i, t := range v.tabs {
		v.tabs[i] = t
	}
	// v.tabs = []string{
	// 	"Skills",
	// 	"Inventory",
	// 	"Ships",
	// 	"Quests",
	// 	"Map",
	// 	"Industry",
	// 	"Credits",
	// }

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

	// Delegate command to current view
	v.currTab, cmd = v.tabs[v.currTabId].Update(msg)
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
				v.currTab.View(),
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
		if v.currTabId == i {
			titleStyle = titleStyle.Copy().Bold(true).Underline(true)
		}
		tabsTitle = append(tabsTitle, titleStyle.Render(tab.Title()))
	}
	return lipgloss.JoinVertical(lipgloss.Left, tabsTitle...)
}

func (v *MainView) setCurrentTab(tab int) {
	v.currTabId = tab
}

func (v MainView) getTabAt(id int) tab.Tab {
	tabs := v.tabs
	return tabs[id]
}

func (v MainView) getPrevTabId() int {
	var currTabId int
	tabs := v.tabs
	currTabId = (v.currTabId - 1) % len(tabs)
	if currTabId < 0 {
		currTabId += len(tabs)
	}

	return currTabId
}

func (v MainView) getNextTabId() int {
	return (v.currTabId + 1) % len(v.tabs)
}
