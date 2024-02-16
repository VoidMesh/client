package view

import (
	"github.com/VoidMesh/client/internal/context"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Ctx *context.ProgramContext
}

type View interface {
	Update(msg tea.Msg) (View, tea.Cmd)
	Render() string
}
