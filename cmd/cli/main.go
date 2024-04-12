package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	rootDir := "D:\\git\\zuehlke\\platformplane\\platformplane"
	// Compile the regular expression pattern
	re := regexp.MustCompile(`//\s?(?P<url>https://\S+\w)\s*version = "(?P<version>[\d.]+)"`)

	columns := []table.Column{
		{Title: "File", Width: 20},
		{Title: "Artifact", Width: 120},
		{Title: "Version", Width: 40},
	}

	var rows []table.Row

	// Call the walk function to traverse all files and subdirectories
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)

			if err != nil {
				return err
			}

			fileContent := string(content)

			match := re.FindStringSubmatch(fileContent)

			if match != nil {
				rows = append(rows, table.Row{filepath.Base(path), match[1], match[2]})
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := model{t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}
