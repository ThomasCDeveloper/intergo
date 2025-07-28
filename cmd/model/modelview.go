package model

import "intergo/cmd/styles"

func (m Model) View() string {
	switch m.mode {
	case tooSmall:
		return m.TooSmallView()
	case boot:
		return m.BootView()
	case menu:
		return m.MenuView()
	case chat:
		return m.ChatView()
	}
	return ""
}

func (m Model) TooSmallView() string {
	msg := " Your terminal is too small "
	return styles.Styles["border"].MarginLeft((m.viewportWidth - len(msg)) / 2).Render(msg)
}

func (m Model) BootView() string {
	return "Booting..."
}

func (m Model) MenuView() string {
	return "Menu"
}

func (m Model) ChatView() string {
	return "Chat"
}
