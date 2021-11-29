package orcid

import (
	"net/http"
)

type History struct {
	Claimed              bool      `json:"claimed,omitempty"`
	CompletionDate       TimeValue `json:"completion-date,omitempty"`
	CreationMethod       string    `json:"creation-method,omitempty"`
	DeactivationDate     TimeValue `json:"deactivation-date,omitempty"`
	LastModifiedDate     TimeValue `json:"last-modified-date,omitempty"`
	SubmissionDate       TimeValue `json:"submission-date,omitempty"`
	Source               *Source   `json:"source,omitempty"`
	VerifiedEmail        bool      `json:"verified-email,omitempty"`
	VerifiedPrimaryEmail bool      `json:"verified-primary-email,omitempty"`
}

type Preferences struct {
	Locale string `json:"locale,omitempty"`
}

// TODO activities-summary
type Record struct {
	History         *History     `json:"history,omitempty"`
	OrcidIdentifier URI          `json:"orcid-identifier,omitempty"`
	Person          *Person      `json:"person,omitempty"`
	Path            string       `json:"path,omitempty"`
	Preferences     *Preferences `json:"preferences,omitempty"`
}

// TODO set data to nil if error?
func (c *Client) Record(orcid string) (*Record, *http.Response, error) {
	data := &Record{}
	res, err := c.get(orcid, data)
	return data, res, err
}
