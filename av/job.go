package av

import (
	"path"
	"time"

	"github.com/gofrs/uuid"
)

// TagTranscodeDefault runs any default transcodes
// TagDolbyVisionTranscode runs Dolby Vision transcodes
// TagDolbyVisionPreprocess runs Dolby Vision pre-processing
// TagDolbyVisionMezzQC runs QC check on the mezzanine
const (
	TagTranscodeDefault      = "transcodeDefault"
	TagDolbyVisionTranscode  = "doViTranscode"
	TagDolbyVisionPreprocess = "doViPreprocess"
	TagDolbyVisionMezzQC     = "doViMezzQC"
	TagDolbyVisionMetadata   = "dolbyVisionMetadata" // inconsistent
)

// Env contains configurations for the environment used while transcoding
type Env struct {
	Cloud       string            `json:"cloud"`
	Region      string            `json:"region"`
	InputAlias  string            `json:"inputAlias"`
	OutputAlias string            `json:"outputAlias"`
	Tags        map[string]string `json:"tags,omitempty"`
}

// Job is a transcoding job
type Job struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Labels []string `json:"labels,omitempty"`

	CreatedAt     time.Time `json:"createdAt"`
	Provider      string    `json:"provider"`
	ProviderJobID string    `json:"providerJobID"`

	Input  File `json:"input"`
	Output Dir  `json:"output"`

	Features Features `json:"features,omitempty"`
	Env      Env      `json:"env"`

	ExtraFiles map[string]string `json:"extraFiles,omitempty"`
}

func (j *Job) Asset(sidecar string) *File {
	loc := j.ExtraFiles[sidecar]
	if loc == "" {
		return nil
	}
	return &File{Name: loc}
}

func (j *Job) Dir() File {
	return File{Name: j.Location("")}
}

func (j Job) Location(file string) string {
	u := j.Output.Location()
	u.Path = path.Join(u.Path, j.rootFolder(), file)
	return u.String()
}
func (j *Job) Abs(f File) File {
	f.Name = j.Location(f.Name)
	return f
}
func (j Job) rootFolder() string {
	if j.Name != "" {
		if _, err := uuid.FromString(j.Name); err == nil {
			return j.Name
		}
	}
	return j.ID
}

// Features is a map whose key is a custom feature name and value is a json string
// representing the corresponding custom feature definition
type Features map[string]interface{}

//Overlays defines all the overlay settings for a Video preset
type Overlays struct {
	Images         []Image   `json:"images,omitempty"`
	TimecodeBurnin *Timecode `json:"timecodeBurnin,omitempty"`
}

//Image defines the image overlay settings
type Image struct {
	URL string `json:"url"`
}

// Timecode settings
type Timecode struct {
	FontSize int    `json:"fontSize,omitempty"`
	Position int    `json:"position,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

// HDR10 configurations and metadata
type HDR10 struct {
	Enabled       bool   `json:"enabled"`
	MaxCLL        int    `json:"maxCLL,omitempty"`
	MaxFALL       int    `json:"maxFALL,omitempty"`
	MasterDisplay string `json:"masterDisplay,omitempty"`
}

// DolbyVision settings
type DolbyVision struct {
	Enabled bool `json:"enabled"`
}
