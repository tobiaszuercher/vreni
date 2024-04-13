package scanner

import (
	"errors"

	"github.com/Masterminds/semver/v3"

	"github.com/tobiaszuercher/vreni/pkg/artifacthub"
	"github.com/tobiaszuercher/vreni/pkg/model"
)

func (s *Scanner) Check(artifacts []*model.Artifact) error {
	var result error

	hub := artifacthub.New(s.Config)

	for _, a := range artifacts {
		versions, err := hub.PackageVersions(a.URL)

		if err != nil {
			result = errors.Join(result, err)
		}

		if len(versions) > 0 {
			v, err := semver.NewVersion(versions[0])

			if err != nil {
				result = errors.Join(result, err)
				continue
			}

			a.AvailableVersion = v
		}
	}

	return result
}
