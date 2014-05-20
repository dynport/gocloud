package drive

import "net/http"

type FilesList struct {
	Q          string
	PageToken  string
	MaxResults string
}

func (f *FilesList) Execute(client *http.Client) (*FilesListResponse, error) {
	req, e := http.NewRequest("GET", "https://www.googleapis.com/drive/v2/files", nil)
	if e != nil {
		return nil, e
	}
	r := &FilesListResponse{}
	e = loadRequest(client, req, r)
	return r, e
}

type FilesListResponse struct {
	Kind          string  `json:"kind"`          // "drive#fileList",
	Etag          string  `json:"etag"`          // etag,
	SelfLink      string  `json:"selfLink"`      // string,
	NextPageToken string  `json:"nextPageToken"` // string,
	NextLink      string  `json:"nextLink"`      // string,
	Items         []*File `json:"items"`
}

type Thumbnail struct {
	Image    string
	MimeType string
}

type File struct {
	Kind                  string              `json:"kind"`           // "drive#file",
	Id                    string              `json:"id"`             // string,
	Etag                  string              `json:"etag"`           // etag,
	SelfLink              string              `json:"selfLink"`       // string,
	WebContentLink        string              `json:"webContentLink"` // string,
	WebViewLink           string              `json:"webViewLink"`    // string,
	AlternateLink         string              `json:"alternateLink"`  // string,
	EmbedLink             string              `json:"embedLink"`      // string,
	OpenWithLinks         map[string]string   `json:"openWithLinks"`
	DefaultOpenWithLink   string              `json:"defaultOpenWithLink"` // string,
	IconLink              string              `json:"iconLink"`            // string,
	ThumbnailLink         string              `json:"thumbnailLink"`       // string,
	Thumbnail             *Thumbnail          `json:"thumbnail"`
	Title                 string              `json:"title"`       // string,
	MimeType              string              `json:"mimeType"`    // string,
	Description           string              `json:"description"` // string,
	Labels                *Labels             `json:"labels"`
	CreatedDate           string              `json:"createdDate"`        // datetime,
	ModifiedDate          string              `json:"modifiedDate"`       // datetime,
	ModifiedByMeDate      string              `json:"modifiedByMeDate"`   // datetime,
	LastViewedByMeDate    string              `json:"lastViewedByMeDate"` // datetime,
	SharedWithMeDate      string              `json:"sharedWithMeDate"`   // datetime,
	Parents               []*Parent           `json:"parents"`
	DownloadUrl           string              `json:"downloadUrl"` // string,
	ExportLinks           interface{}         `json:"exportLinks"`
	IndexableText         map[string]string   `json:"indexableText"`
	UserPermission        *Permissions        `json:"userPermission"`   // permissions Resource,
	OriginalFilename      string              `json:"originalFilename"` // string,
	FileExtension         string              `json:"fileExtension"`    // string,
	Md5Checksum           string              `json:"md5Checksum"`      // string,
	FileSize              string              `json:"fileSize"`         // long,
	QuotaBytesUsed        string              `json:"quotaBytesUsed"`   // long,
	OwnerNames            []string            `json:"ownerNames"`
	Owners                []*User             `json:"owners"`
	LastModifyingUserName string              `json:"lastModifyingUserName"` // string,
	LastModifyingUser     *User               `json:"lastModifyingUser"`
	Editable              bool                `json:"editable"`          // boolean,
	Copyable              bool                `json:"copyable"`          // boolean,
	WritersCanShare       bool                `json:"writersCanShare"`   // boolean,
	Shared                bool                `json:"shared"`            // boolean,
	ExplicitlyTrashed     bool                `json:"explicitlyTrashed"` // boolean,
	AppDataContents       bool                `json:"appDataContents"`   // boolean,
	HeadRevisionId        string              `json:"headRevisionId"`    // string,
	Properties            []*interface{}      `json:"properties"`
	ImageMediaMetadata    *ImageMediaMetadata `json:"imagemediametadata"`
}

type Parent struct {
	Kind       string `json:"kind"`       // "drive#parentReference",
	Id         string `json:"id"`         // string,
	SelfLink   string `json:"selfLink"`   // string,
	ParentLink string `json:"parentLink"` // string,
	IsRoot     bool   `json:"isRoot"`     // boolean
}

type Permissions struct {
	Kind            string   `json:"kind"`         // "drive#permission",
	Etag            string   `json:"etag"`         // etag,
	Id              string   `json:"id"`           // string,
	SelfLink        string   `json:"selfLink"`     // string,
	Name            string   `json:"name"`         // string,
	EmailAddress    string   `json:"emailAddress"` // string,
	Domain          string   `json:"domain"`       // string,
	Role            string   `json:"role"`         // string,
	AdditionalRoles []string `json:"additionalRoles"`
	Type            string   `json:"type"`      // string,
	Value           string   `json:"value"`     // string,
	AuthKey         string   `json:"authKey"`   // string,
	WithLink        bool     `json:"withLink"`  // boolean,
	PhotoLink       string   `json:"photoLink"` // string
}

type Labels struct {
	Starred    bool `json:"starred"`    // boolean,
	Hidden     bool `json:"hidden"`     // boolean,
	Trashed    bool `json:"trashed"`    // boolean,
	Restricted bool `json:"restricted"` // boolean,
	Viewed     bool `json:"viewed"`     // boolean
}

type ImageMediaMetadata struct {
	Width            string    `json:"width"`    // integer,
	Height           string    `json:"height"`   // integer,
	Rotation         string    `json:"rotation"` // integer,
	Location         *Location `json:"location"`
	Date             string    `json:"date"`             // string,
	CameraMake       string    `json:"cameraMake"`       // string,
	CameraModel      string    `json:"cameraModel"`      // string,
	ExposureTime     string    `json:"exposureTime"`     // float,
	Aperture         string    `json:"aperture"`         // float,
	FlashUsed        string    `json:"flashUsed"`        // boolean,
	FocalLength      string    `json:"focalLength"`      // float,
	IsoSpeed         string    `json:"isoSpeed"`         // integer,
	MeteringMode     string    `json:"meteringMode"`     // string,
	Sensor           string    `json:"sensor"`           // string,
	ExposureMode     string    `json:"exposureMode"`     // string,
	ColorSpace       string    `json:"colorSpace"`       // string,
	WhiteBalance     string    `json:"whiteBalance"`     // string,
	ExposureBias     string    `json:"exposureBias"`     // float,
	MaxApertureValue string    `json:"maxApertureValue"` // float,
	SubjectDistance  string    `json:"subjectDistance"`  // integer,
	Lens             string    `json:"lens"`             // string
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
