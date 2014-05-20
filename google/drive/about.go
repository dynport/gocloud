package drive

import "net/http"

type About struct {
}

func (a *About) Execute(client *http.Client) (*AboutResponse, error) {
	req, e := http.NewRequest("GET", "https://www.googleapis.com/drive/v2/about", nil)
	if e != nil {
		return nil, e
	}
	r := &AboutResponse{}
	e = loadRequest(client, req, r)
	return r, e
}

type AboutResponse struct {
	Kind                    string        `json:"kind"`     // "drive#about",
	Etag                    string        `json:"etag"`     // etag,
	SelfLink                string        `json:"selfLink"` // string,
	Name                    string        `json:"name"`     // string,
	User                    *User         `json:"user"`
	QuotaBytesTotal         string        `json:"quotaBytesTotal"`         // long,
	QuotaBytesUsed          string        `json:"quotaBytesUsed"`          // long,
	QuotaBytesUsedAggregate string        `json:"quotaBytesUsedAggregate"` // long,
	QuotaBytesUsedInTrash   string        `json:"quotaBytesUsedInTrash"`   // long,
	LargestChangeId         string        `json:"largestChangeId"`         // long,
	RemainingChangeIds      string        `json:"remainingChangeIds"`      // long,
	RootFolderId            string        `json:"rootFolderId"`            // string,
	DomainSharingPolicy     string        `json:"domainSharingPolicy"`     // string,
	PermissionId            string        `json:"permissionId"`            // string,
	IsCurrentAppInstalled   bool          `json:"isCurrentAppInstalled"`
	ImportFormats           []*Format     `json:"importFormats"`
	ExportFormats           []*Format     `json:"exportFormats"`
	AdditionalRoleInfos     []*RoleInfo   `json:"additionalRoleInfo"`
	Features                []*Feature    `json:"features"`
	MaxUploadSizes          []*UploadSize `json:"maxUploadSizes"`
}

type RoleInfo struct {
	Type     string     `json:"type"`
	RoleSets []*RoleSet `json:"roleSets"`
}

type RoleSet struct {
	PrimaryRole     string   `json:"primaryRole"`
	AdditionalRoles []string `json:"additionalRoles"`
}

type Feature struct {
	FeatureName string  `json:"featureName"`
	FeatureRate float64 `json:"featureRate"`
}

type UploadSize struct {
	Type string `json:"type"`
	Size string `json:"size"`
}

type Format struct {
	Source  string   `json:"source"`
	Targets []string `json:"targets"`
}

type Picture struct {
	Url string `json:"url"` // string
}

type User struct {
	Kind                string   `json:"kind"`        // "drive#user",
	DisplayName         string   `json:"displayName"` // string,
	Picture             *Picture `json:"picture"`
	IsAuthenticatedUser bool     `json:"isAuthenticatedUser"` // boolean,
	PermissionId        string   `json:"permissionId"`        // string
}

type Drive struct {
	*http.Client
}
