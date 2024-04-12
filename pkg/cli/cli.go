package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/tobiaszuercher/vervet/pkg/model"
)

func ArtifactTable(artifacts []*model.Artifact) {
	columns := []table.Column{
		{Title: "File", Width: 20},
		{Title: "Artifact", Width: 100},
		{Title: "Version", Width: 40},
		{Title: "Latest", Width: 40},
	}

	rows := generateRowsFromArtifacts(artifacts)

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

	m := TableModel{t, artifacts}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

type TableModel struct {
	table table.Model

	data []*model.Artifact
}

func (m TableModel) Init() tea.Cmd { return nil }

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case []*model.Artifact:
		m.data = msg

		m.table.SetRows(generateRowsFromArtifacts(msg))
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m TableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func generateRowsFromArtifacts(artifacts []*model.Artifact) []table.Row {
	rows := make([]table.Row, 0)

	for _, a := range artifacts {
		rows = append(rows, table.Row{
			a.File,
			a.URL,
			a.Version,
			a.AvailableVersion,
		})
	}

	return rows
}
