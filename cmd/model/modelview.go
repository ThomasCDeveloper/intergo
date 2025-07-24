package model

import "intergo/cmd/styles"

func (m Model) View() string {
	switch m.mode {
	case tooSmall:
		return styles.Render("Window too small", "highlighted")
	case boot:
		return "boot"
	case menu:
		return "menu"
	case chat:
		return "chat"
	}
	return ""
}
