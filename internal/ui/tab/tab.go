package tab

import tea "github.com/charmbracelet/bubbletea"

type TabBase struct {
	Name string
}

type Tab interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Tab, tea.Cmd)
	View() string
	Title() string
}

type Model struct {
	Name string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Tab, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Example Tab View"
}

func (m Model) Title() string {
	return m.Name
}
