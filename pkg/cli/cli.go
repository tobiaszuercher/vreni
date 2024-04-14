package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/tobiaszuercher/vreni/pkg/model"
)

const (
	purple = lipgloss.Color("99")
	white  = lipgloss.Color("255")
)

func List(artifacts []*model.Artifact) {
	re := lipgloss.NewRenderer(os.Stdout)
	rows := createRows(artifacts, re)

	var (
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Padding(0, 2)
		CellStyle   = re.NewStyle().Padding(0, 2)
		OddRowStyle = CellStyle.Copy().Foreground(white)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return HeaderStyle
			}

			switch col {
			case 7:
				return lipgloss.NewStyle().Background(lipgloss.Color("99"))
			}

			return OddRowStyle
		}).
		Headers("FILE", "ARTIFACT", "INSTALLED", "AVAILABLE").
		Rows(rows...)

	fmt.Println(t)
}

func createRows(artifacts []*model.Artifact, re *lipgloss.Renderer) [][]string {
	var rows [][]string

	for _, a := range artifacts {
		rows = append(rows, []string{filepath.Base(a.File), a.URL, a.Version.String(), a.ColoredDiff()})
	}

	return rows
}

func Prompt() bool {
	var ok bool

	fmt.Println()

	confirm := huh.NewConfirm().
		Title("Do you want to update all deps?").
		Affirmative("Yes").
		Negative("No").
		Value(&ok)

	huh.NewForm(huh.NewGroup(confirm)).Run()

	if ok {
		fmt.Println("Updating deps...")
	} else {
		fmt.Println("No deps updated.")
	}

	return ok
}
