package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func LaunchProgram() {
	program := tea.NewProgram(initialiseModel())
	if _, err := program.Run(); err != nil {
		fmt.Println("Failed to launch the program:", err)
	}
}
