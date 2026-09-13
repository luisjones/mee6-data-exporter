package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func inputView(m model) string {
	statusColor := colourInvalidID
	statusText := "Invalid ID"
	if m.isValidDiscordGuildID() {
		statusColor = colourValidID
		statusText = "Valid ID"
	}

	status := lipgloss.NewStyle().Foreground(statusColor).Render(statusText)
	title := lipgloss.NewStyle().Bold(true).Foreground(colourTitle).Render("MEE6 Exporter")

	body := "This program will export MEE6 server statistics.\n\n  " +
		"To get started, enter your Discord Server ID.\n  " +
		"You can find this by right clicking your server\n  " +
		"icon in Discord and clicking \"Copy Server ID\".\n\n  " +
		"%s\n  %s"

	return fmt.Sprintf("%s\n\n  "+body, title, m.TextInput.View(), status) + "\n"
}

func spinnerView(m model) string {
	if m.Finished {
		return "\n\n" + m.CurrentStatus
	}
	label := m.Spinner.View() + m.CurrentStatus
	return "\n\n" + label
}
