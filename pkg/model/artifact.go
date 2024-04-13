package model

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/charmbracelet/lipgloss"
)

type Artifact struct {
	File string
	URL  string

	Version          *semver.Version
	AvailableVersion *semver.Version

	VersionsBetween int
}

type UpdateLevel int

const (
	NoUpdate UpdateLevel = 0
	Major    UpdateLevel = 1
	Patch    UpdateLevel = 2
	Minor    UpdateLevel = 3
)

func (a *Artifact) HasUpdate() bool {
	return a.Version != a.AvailableVersion
}

func (a *Artifact) UpgradeLevel() UpdateLevel {
	if a.Version.Major() != a.AvailableVersion.Major() {
		return Major
	}

	if a.Version.Minor() != a.AvailableVersion.Minor() {
		return Patch
	}

	if a.Version.Patch() != a.AvailableVersion.Patch() {
		return Minor
	}

	return NoUpdate
}

func (a *Artifact) RenderDiff() string {
	red := lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Inline(true)
	green := lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Inline(true)
	blue := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Inline(true)
	white := lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Inline(true)

	if a.UpgradeLevel() == Major {
		return red.Render(a.AvailableVersion.String())
	}

	if a.UpgradeLevel() == Patch {
		split := strings.Split(a.AvailableVersion.String(), ".")

		unchanged := white.Render(split[0] + ".")
		updated := blue.Render(split[1] + "." + split[2])

		return white.Render(unchanged) + blue.Render(updated)
	}

	if a.UpgradeLevel() == Minor {
		split := strings.Split(a.AvailableVersion.String(), ".")

		unchanged := split[0] + "." + split[1] + "."
		updated := split[2]

		return white.Render(unchanged) + green.Render(updated)
	}

	return white.Render(a.AvailableVersion.String())
}
