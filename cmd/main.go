package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yatiac/go-todo-cli/internal/tui"
)

const (
	width       = 96
	columnWidth = 30
)

type command struct {
	disabled bool
	name     string
}

type model struct {
	stateDescription string
	stateStatus      tui.StatusBarState
	commands         []command
	cursor           int
}

func initialModel() model {
	return model{
		stateDescription: "Initializing...",
		commands: []command{
			{name: "Add Todo", disabled: false},
			{name: "List Todos", disabled: false},
			{name: "Complete Todo", disabled: true},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	doc := &strings.Builder{}
	// Footer with commands
	tui.RenderTitleRow(width, doc, tui.TitleRowProps{Title: "My Todo App"})
	doc.WriteString("\n\n")
	doc.WriteString("Press q to quit.\n\n")
	tui.RenderStatusBar(doc, tui.NewStatusBarProps(&tui.StatusBarProps{
		Description: m.stateDescription,
		User:        "NONE",
		StatusState: tui.StatusBarStateBlue,
		Width:       width,
	}))
	return doc.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
