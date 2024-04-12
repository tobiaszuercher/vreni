package scanner

import (
	"errors"

	"github.com/tobiaszuercher/vervet/pkg/artifacthub"
	"github.com/tobiaszuercher/vervet/pkg/model"
)

func (s *Scanner) Check(artifacts []*model.Artifact) ([]*model.Artifact, error) {
	var result error

	hub := artifacthub.New(s.Config)

	for _, a := range artifacts {
		versions, err := hub.PackageVersions(a.URL)

		if err != nil {
			result = errors.Join(result, err)
		}

		if len(versions) > 0 {
			a.AvailableVersion = versions[0]
		}
	}

	return artifacts, result
}
