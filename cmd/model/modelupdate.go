package model

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height

		if m.viewportWidth < m.appWidth && m.mode != tooSmall {
			m.lastMode = m.mode
			m.mode = tooSmall
		}

		if m.viewportWidth >= m.appWidth && m.mode == tooSmall {
			m.mode = m.lastMode
		}
	}

	if m.mode == tooSmall {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "b":
			m.mode = boot
		case "m":
			m.mode = menu
		case "c":
			m.mode = chat
		}
	}

	return m, nil
}
