package view

import (
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/client/internal/game"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type CharacterView struct {
	model      tea.Model
	form       *huh.Form
	game       game.Game
	characters []*character.Character
}

func NewCharacterView(g game.Game) tea.Model {
	v := CharacterView{game: g}

	var charOptions = []huh.Option[string]{}

	for _, c := range v.characters {
		charOptions = append(charOptions, huh.NewOption(c.Name, c.Id))
	}

	v.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your character").
				Options(charOptions...),
		),
	)

	return v
}

func (v CharacterView) Init() tea.Cmd {
	return v.form.Init()
}

func (v CharacterView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Process the form
	form, cmd := v.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		v.form = f
		cmds = append(cmds, cmd)
	}

	if v.form.State == huh.StateCompleted {
		// Move to the main view once the form is completed
		view := NewMainView(&v.model)
		return view, cmd
	}

	return v, tea.Batch(cmds...)
}

func (v CharacterView) View() string {
	switch v.form.State {
	case huh.StateCompleted:
		return lipgloss.NewStyle().Render("State completed!")
	default:
		err := v.form.Run()
		if err != nil {
			return lipgloss.NewStyle().Render(err.Error())
		}

		return lipgloss.NewStyle().Render("State not completed.")
	}
}
