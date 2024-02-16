package internal

import (
	"time"

	"github.com/VoidMesh/client/internal/constants"
	"github.com/VoidMesh/client/internal/context"
	"github.com/VoidMesh/client/internal/ui/view"
	"github.com/VoidMesh/client/internal/utils"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Use gRPC streaming to receive updates
type tickMsg time.Time

type Model struct {
	Tick     constants.Tick
	keys     utils.KeyMap
	err      error
	currView view.View
	ctx      *context.ProgramContext
}

func NewModel() *Model {
	m := Model{
		keys: utils.Keys,
		ctx:  &context.ProgramContext{},
		Tick: constants.Tick{
			Duration: constants.TickDuration,
		},
	}

	m.currView = view.NewAuthView(m.ctx)

	return &m
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			cmds = append(cmds, tea.Quit)
		}

	case tickMsg:
		cmd = tea.Batch(tickCmd(m))

	case initMsg:
		cmd = tea.Batch(tickCmd(m))

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case view.View:
		m.currView = msg

	case errMsg:
		m.err = msg
	}

	cmds = append(cmds, cmd)

	// Delegate command to current view
	m.currView, cmd = m.currView.Update(msg)
	cmds = append(cmds, cmd)

	// Return updated model and any new commands
	return m, tea.Batch(cmds...)
}

func tickCmd(m *Model) tea.Cmd {
	return tea.Every(m.Tick.Duration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	m.ctx.Width = msg.Width
	m.ctx.Height = msg.Height
}

func (m Model) View() string {
	return m.currView.Render()
}
