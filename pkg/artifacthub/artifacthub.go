package artifacthub

import (
	"errors"
	"log"
	"net/url"

	"github.com/imroc/req/v3"

	"github.com/tobiaszuercher/vervet/config"
)

type Artifacthub struct {
	*config.Config
}

func New(cfg *config.Config) *Artifacthub {
	return &Artifacthub{
		cfg,
	}
}

func (a *Artifacthub) PackageVersions(artifact string) ([]string, error) {
	client := req.C()

	var result PackageDetail

	u, err := url.Parse(artifact)

	if err != nil {
		return nil, nil
	}

	resp, err := client.R().
		SetHeader("X-API-KEY-ID", a.APIKeyID).
		SetHeader("X-API-KEY-SECRET", a.APISecret).
		SetSuccessResult(&result).
		Get("https://artifacthub.io/api/v1" + u.Path)

	if err != nil {
		log.Println(resp.Dump())
		log.Fatal(err)
	}

	if resp.IsSuccessState() {
		var versions []string

		for _, v := range result.AvailableVersions {
			versions = append(versions, v.Version)
		}

		return versions, nil
	}

	return nil, errors.New("failed to fetch package versions")
}
