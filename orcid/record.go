package orcid

import (
	"net/http"
)

type History struct {
	Claimed        *bool     `json:"claimed,omitempty"`
	CompletionDate *IntValue `json:"completion-date,omitempty"`
	// TODO define enum
	CreationMethod       *string   `json:"creation-method,omitempty"`
	DeactivationDate     *IntValue `json:"deactivation-date,omitempty"`
	LastModifiedDate     *IntValue `json:"last-modified-date,omitempty"`
	SubmissionDate       *IntValue `json:"submission-date,omitempty"`
	Source               *Source   `json:"source,omitempty"`
	VerifiedEmail        *bool     `json:"verified-email,omitempty"`
	VerifiedPrimaryEmail *bool     `json:"verified-primary-email,omitempty"`
}

type Preferences struct {
	// TODO define enum
	Locale string `json:"locale,omitempty"`
}

// TODO activities-summary
type Record struct {
	History         *History     `json:history,omitempty"`
	OrcidIdentifier *URI         `json:"orcid-identifier,omitempty"`
	Person          *Person      `json:"person,omitempty"`
	Path            *string      `json:"path,omitempty"`
	Preferences     *Preferences `json:"preferences,omitempty"`
}

func (c *Client) Record(orcid string) (*Record, *http.Response, error) {
	data := new(Record)
	res, err := c.get(orcid, data)
	return data, res, err
}
