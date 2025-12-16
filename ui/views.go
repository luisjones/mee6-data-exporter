package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func inputView(m model) string {
	var style lipgloss.Style

	if m.isValidDiscordGuildID() {
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#32cd32"))
		return fmt.Sprintf(
			"Discord Server ID\n\n%s\n%s",
			m.TextInput.View(),
			style.Render("Valid ID"),
		) + "\n"
	} else {
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
		return fmt.Sprintf(
			"Discord Server ID\n\n%s\n%s",
			m.TextInput.View(),
			style.Render("Invalid ID"),
		) + "\n"
	}
}

func spinnerView(m model) string {
	if m.Finished {
		return "\n\n" + m.CurrentStatus
	}
	label := m.Spinner.View() + m.CurrentStatus
	return "\n\n" + label
}
