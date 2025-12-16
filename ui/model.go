package ui

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"mee6xport/db"
	"mee6xport/mee6"
	"mee6xport/ui/components"
	"regexp"
	"strconv"
)

// This holds the application state
type model struct {
	TextInput     textinput.Model
	Spinner       spinner.Model
	InputEntered  bool
	Quitting      bool
	CurrentStatus string

	StartGenerating bool
	Pages           []mee6.Response
	DB              *sql.DB
	Tx              *sql.Tx

	Finished bool
}

type CompletionMsg struct{}

func (m model) CrawlAndInsert() tea.Cmd {
	return func() tea.Msg {
		guildID, _ := strconv.Atoi(m.TextInput.Value())
		database, tx := db.PrepareDB()
		defer database.Close()
		pages, _ := mee6.CrawlGuild(guildID)
		for _, page := range pages {
			page.Insert(tx)
		}
		db.CommitTransaction(tx)
		return CompletionMsg{}
	}
}

// Creates a new model{} structure, using default config
func initialiseModel() model {
	return model{
		TextInput:       components.TextInput(),
		Spinner:         components.Spinner(),
		InputEntered:    false,
		Finished:        false,
		Quitting:        false,
		CurrentStatus:   "",
		StartGenerating: false,
		Pages:           []mee6.Response{},
	}
}

func (m model) Init() tea.Cmd {
	return m.Spinner.Tick
}

// Update processes events and updates the model state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	// Always update spinner first
	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if key == "ctrl+c" || key == "esc" {
			m.Quitting = true
			return m, tea.Quit
		}
		// Only update text input if we haven't entered yet
		if !m.InputEntered {
			m.TextInput, _ = m.TextInput.Update(msg)
		}

		// Handle enter key
		if key == "enter" && !m.InputEntered && m.isValidDiscordGuildID() {
			m.InputEntered = true
			m.StartGenerating = true
			m.CurrentStatus = "crawling guild data..."
			cmds = append(cmds, m.CrawlAndInsert())
			return m, tea.Batch(cmds...)
		}

	case CompletionMsg:
		m.Finished = true
		m.CurrentStatus = "Database saved successfully!"
		return m, tea.Batch(cmds...)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  Written by Luis / github.com/luisjones\n\n"
	}

	s = inputView(m)

	if m.InputEntered {
		s = spinnerView(m)
	}

	return indent.String(fmt.Sprintf("\n%s\n\n", s), 2)
}

func (m model) isValidDiscordGuildID() bool {
	// Checks for digits with a length of 17-19 characters.
	regex, _ := regexp.Compile(`\d{17,19}`)
	return regex.MatchString(m.TextInput.Value())
}
