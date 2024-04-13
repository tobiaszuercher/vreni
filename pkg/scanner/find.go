package scanner

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"

	"github.com/tobiaszuercher/vreni/pkg/model"
)

func (s *Scanner) Find(dir string) ([]*model.Artifact, error) {
	re := regexp.MustCompile(`//\s?(https://\S+\w)\s*version = "([\d.]+)"`)

	var result []*model.Artifact

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)

			if err != nil {
				return nil
			}

			fileContent := string(content)

			match := re.FindStringSubmatch(fileContent)

			if match != nil {
				// TODO: there are also hub.docker.com
				if !strings.Contains(match[1], "https://artifacthub.io") {
					return nil
				}

				v, err := semver.NewVersion(match[2])

				if err != nil {
					return err
				}

				result = append(result, &model.Artifact{
					File:    path,
					URL:     match[1],
					Version: v,
				})
			}
		}

		return nil
	})

	return result, err
}
