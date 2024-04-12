package artifacthub

import (
	"fmt"
	"os"
	"testing"
)

func TestArtifacthub(t *testing.T) {
	a := New(os.Getenv("API_KEY_ID"), os.Getenv("API_KEY_SECRET"))

	versions, err := a.PackageVersions("packages/helm/grafana/loki")

	if err != nil {
		t.Error(err)
	}

	if len(versions) < 1 {
		t.Error("no versions found")
	}

	fmt.Printf("%v", versions)
}
