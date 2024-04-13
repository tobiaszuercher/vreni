package artifacthub

import (
	"fmt"
	"testing"

	"github.com/tobiaszuercher/vervet/config"
)

func TestArtifacthub(t *testing.T) {
	a := New(config.FromEnvironment())

	versions, err := a.PackageVersions("packages/helm/grafana/loki")

	if err != nil {
		t.Error(err)
	}

	if len(versions) < 1 {
		t.Error("no versions found")
	}

	fmt.Printf("%v", versions)
}
