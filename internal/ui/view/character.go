package view

import (
	"context"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/client/internal/game"
	"github.com/VoidMesh/client/internal/program_context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type CharacterView struct {
	model tea.Model
	form  *huh.Form
	game  game.Game
	ctx   *program_context.Ctx
}

func NewCharacterView(ctx *program_context.Ctx, g game.Game) tea.Model {
	v := CharacterView{ctx: ctx, game: g}

	resp, err := g.Services.Character.List(context.TODO(), &character.ListRequest{
		AccountId: g.Account.Id,
	})

	if err != nil {
		log.Fatal(err)
	}

	v.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[*character.Character]().
				Key("character").
				Title("Choose your character").
				Options(huh.NewOptions(resp.Characters...)...),
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
		v.game.Character = v.form.Get("character").(*character.Character)
		view := NewMainView(v.ctx, v.game)
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
