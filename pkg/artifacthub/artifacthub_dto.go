package artifacthub

type PackageDetail struct {
	PackageId      string   `json:"package_id"`
	Name           string   `json:"name"`
	NormalizedName string   `json:"normalized_name"`
	Category       int      `json:"category"`
	IsOperator     bool     `json:"is_operator"`
	Description    string   `json:"description"`
	LogoImageId    string   `json:"logo_image_id"`
	Keywords       []string `json:"keywords"`
	HomeUrl        string   `json:"home_url"`
	Readme         string   `json:"readme"`
	Links          []struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"links"`

	SecurityReportSummary struct {
		Low      int `json:"low"`
		High     int `json:"high"`
		Medium   int `json:"medium"`
		Unknown  int `json:"unknown"`
		Critical int `json:"critical"`
	} `json:"security_report_summary"`

	SecurityReportCreatedAt int `json:"security_report_created_at"`

	Data struct {
		Type         string `json:"type"`
		ApiVersion   string `json:"apiVersion"`
		KubeVersion  string `json:"kubeVersion"`
		Dependencies []struct {
			Name       string `json:"name"`
			Version    string `json:"version"`
			Repository string `json:"repository"`
		} `json:"dependencies"`
	} `json:"data"`

	Version string `json:"version"`

	AvailableVersions []struct {
		Version                 string `json:"version"`
		ContainsSecurityUpdates bool   `json:"contains_security_updates"`
		Prerelease              bool   `json:"prerelease"`
		Ts                      int    `json:"ts"`
	} `json:"available_versions"`

	AppVersion              string `json:"app_version"`
	Digest                  string `json:"digest"`
	Deprecated              bool   `json:"deprecated"`
	ContainsSecurityUpdates bool   `json:"contains_security_updates"`
	Prerelease              bool   `json:"prerelease"`
	Signed                  bool   `json:"signed"`
	ContentUrl              string `json:"content_url"`
	ContainersImages        []struct {
		Name        string `json:"name"`
		Image       string `json:"image"`
		Whitelisted bool   `json:"whitelisted"`
	} `json:"containers_images"`
	AllContainersImagesWhitelisted bool `json:"all_containers_images_whitelisted"`
	HasValuesSchema                bool `json:"has_values_schema"`
	HasChangelog                   bool `json:"has_changelog"`
	Ts                             int  `json:"ts"`

	Repository struct {
		RepositoryId            string `json:"repository_id"`
		Name                    string `json:"name"`
		DisplayName             string `json:"display_name"`
		Url                     string `json:"url"`
		Private                 bool   `json:"private"`
		Kind                    int    `json:"kind"`
		VerifiedPublisher       bool   `json:"verified_publisher"`
		Official                bool   `json:"official"`
		ScannerDisabled         bool   `json:"scanner_disabled"`
		OrganizationName        string `json:"organization_name"`
		OrganizationDisplayName string `json:"organization_display_name"`
	} `json:"repository"`

	Stats struct {
		Subscriptions int `json:"subscriptions"`
		Webhooks      int `json:"webhooks"`
	} `json:"stats"`

	ProductionOrganizationsCount int `json:"production_organizations_count"`
}
