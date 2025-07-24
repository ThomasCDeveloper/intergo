package main

import (
	"fmt"
	"intergo/cmd/model"
	"intergo/cmd/styles"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	styles.InitStyles()
}

func main() {
	var m model.Model = model.DefaultModel()

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
