package tab

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Skills struct {
	TabBase
}

func NewSkillsTab() Skills {
	return Skills{TabBase{Name: "Skills"}}
}

func (s Skills) Init() tea.Cmd {
	return nil
}

func (s Skills) Update(msg tea.Msg) (Tab, tea.Cmd) {
	return s, nil
}

func (s Skills) View() string {
	return "Skills tab"
}

func (s Skills) Title() string {
	return s.TabBase.Name
}
