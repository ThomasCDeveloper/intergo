package main

import (
	"fmt"
	"intergo/cmd/model"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var m model.Model

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
