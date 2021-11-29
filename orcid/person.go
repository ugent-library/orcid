package orcid

import (
	"fmt"
	"net/http"
)

type Person struct {
	Addresses           *Addresses           `json:"addresses,omitempty"`
	Biography           *Biography           `json:"biography,omitempty"`
	Emails              *Emails              `json:"emails,omitempty"`
	ExternalIdentifiers *ExternalIdentifiers `json:"external-identifiers,omitempty"`
	Keywords            *Keywords            `json:"keywords,omitempty"`
	LastModifiedDate    TimeValue            `json:"last-modified-date,omitempty"`
	Name                *Name                `json:"name,omitempty"`
	OtherNames          *OtherNames          `json:"other-names,omitempty"`
	Path                string               `json:"path,omitempty"`
	ResearcherURLs      *ResearcherURLs      `json:"researcher-urls,omitempty"`
}

func (c *Client) Person(orcid string) (*Person, *http.Response, error) {
	data := &Person{}
	path := fmt.Sprintf("%s/person", orcid)
	res, err := c.get(path, data)
	return data, res, err
}
