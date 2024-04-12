package scanner

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/tobiaszuercher/vervet/pkg/model"
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
				return err
			}

			fileContent := string(content)

			match := re.FindStringSubmatch(fileContent)

			if match != nil {
				// TODO: there are also hub.docker.com
				if !strings.Contains(match[1], "https://artifacthub.io") {
					return nil
				}

				result = append(result, &model.Artifact{
					File:    filepath.Base(path),
					URL:     match[1],
					Version: match[2],
				})
			}
		}

		return nil
	})

	return result, err
}
