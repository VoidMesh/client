package tab

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Inventory struct {
	TabBase
}

func NewInventoryTab() Inventory {
	return Inventory{TabBase{Name: "Inventory"}}
}

func (i Inventory) Init() tea.Cmd {
	return nil
}

func (i Inventory) Update(msg tea.Msg) (Tab, tea.Cmd) {
	return i, nil
}

func (i Inventory) View() string {
	return "Inventory tab"
}

func (i Inventory) Title() string {
	return i.TabBase.Name
}
