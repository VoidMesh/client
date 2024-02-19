package view

import (
	"context"
	"log"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/client/internal/game"
	"github.com/VoidMesh/client/internal/program_context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type AuthView struct {
	form *huh.Form
	game game.Game
	ctx  *program_context.Ctx
}

func NewAuthView(ctx *program_context.Ctx, g game.Game) tea.Model {
	v := AuthView{ctx: ctx, game: g}

	v.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("email").
				Title("Email").
				Description("Your email").
				Placeholder("void-mesh@example.com"),
		),
	)

	return v
}

func (v AuthView) Init() tea.Cmd {
	return v.form.Init()
}

func (v AuthView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Process the form
	form, cmd := v.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		v.form = f
		cmds = append(cmds, cmd)
	}

	if v.form.State == huh.StateCompleted {
		resp, err := v.game.Services.Account.Authenticate(context.TODO(), &account.AuthenticateRequest{
			Email: v.form.GetString("email"),
		})

		if err != nil {
			log.Fatal(err)
		}

		// Move to the pick character view once the form is completed
		v.game.Account = resp.Account
		view := NewCharacterView(v.ctx, v.game)
		return view, cmd
	}

	return v, tea.Batch(cmds...)
}

func (v AuthView) View() string {
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
