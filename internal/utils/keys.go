package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	PrevTab key.Binding
	NextTab key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevTab, k.NextTab},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap{
	PrevTab: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "prev tab"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "next tab"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
