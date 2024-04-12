package model

type Artifact struct {
	File string
	URL  string

	Version          string
	AvailableVersion string

	VersionsBetween int

	HasUpdate bool
}
