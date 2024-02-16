package view

import (
	"github.com/VoidMesh/client/internal/context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	statusNormal state = iota
	stateDone
)

type AuthView struct {
	state state
	form  *huh.Form
	width int
	View  *Model
}

func NewAuthView(ctx *context.ProgramContext) AuthView {
	m := AuthView{
		View: &Model{
			Ctx: ctx,
		},
	}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Description("Your username").
				Placeholder("johndoe"),

			huh.NewInput().
				Title("Password").
				Description("Your password").
				Placeholder("password").
				Password(true),
		),
	)

	return m
}

func (v AuthView) Init() tea.Cmd {
	return v.form.Init()
}

func (v AuthView) Update(msg tea.Msg) (View, tea.Cmd) {
	var cmds []tea.Cmd

	// Process the form
	form, cmd := v.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		v.form = f
		cmds = append(cmds, cmd)
	}

	if v.form.State == huh.StateCompleted {
		// Move to the main view once the form is completed
		view := MainView{View: &Model{}}
		return view, cmd
	}

	return v, tea.Batch(cmds...)
}

func (v AuthView) Render() string {
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
