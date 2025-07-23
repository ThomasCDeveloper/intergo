package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	Boot Mode = iota
	Menu
	Chat
)

type Model struct {
	mode Mode
}

// TODO: breakdown into diffrent files
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}

func (m Model) Init() tea.Cmd {
	return nil
}
