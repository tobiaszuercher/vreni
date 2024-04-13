package scanner

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/tobiaszuercher/vreni/pkg/model"
)

func Update(artifacts []*model.Artifact) error {
	var result error

	re := regexp.MustCompile(`//\s?https://\S+\w\s*version = "([\d.]+)"`)

	for _, a := range artifacts {
		content, err := os.ReadFile(a.File)

		if err != nil {
			result = errors.Join(result, err)
		}

		updated := re.ReplaceAllStringFunc(string(content), func(match string) string {
			matches := re.FindStringSubmatch(match)

			return strings.Replace(matches[0], matches[1], a.AvailableVersion.String(), 1)
		})

		if err := os.WriteFile(a.File, []byte(updated), 0644); err != nil {
			return err
		}
	}

	return nil
}
