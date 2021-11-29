package orcid

import (
	"fmt"
	"net/http"
)

type PersonalDetails struct {
	Biography        *Biography  `json:"biography,omitempty"`
	LastModifiedDate TimeValue   `json:"last-modified-date,omitempty"`
	Name             *Name       `json:"name,omitempty"`
	OtherNames       *OtherNames `json:"other-names,omitempty"`
	Path             string      `json:"path,omitempty"`
}

func (c *Client) PersonalDetails(orcid string) (*PersonalDetails, *http.Response, error) {
	data := &PersonalDetails{}
	path := fmt.Sprintf("%s/personal-details", orcid)
	res, err := c.get(path, data)
	return data, res, err
}
