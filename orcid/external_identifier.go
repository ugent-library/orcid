package orcid

import (
	"fmt"
	"net/http"
)

type ExternalIdentifier struct {
	CreatedDate  TimeValue `json:"created-date,omitempty"`
	DisplayIndex int       `json:"display-index,omitempty"`
	// swagger docs mistakenly say object
	ExternalIDRelationship string      `json:"external-id-relationship,omitempty"`
	ExternalIDType         string      `json:"external-id-type,omitempty"`
	ExternalIdValue        string      `json:"external-id-value,omitempty"`
	ExternalIDUrl          StringValue `json:"external-id-url,omitempty"`
	LastModifiedDate       TimeValue   `json:"last-modified-date,omitempty"`
	Path                   string      `json:"path,omitempty"`
	PutCode                int         `json:"put-code,omitempty"`
	Source                 *Source     `json:"source,omitempty"`
	Visibility             string      `json:"visibility,omitempty"`
}

type ExternalIdentifiers struct {
	ExternalIdentifier []ExternalIdentifier `json:"external-identifier,omitempty"`
	LastModifiedDate   TimeValue            `json:"last-modified-date,omitempty"`
	Path               string               `json:"path,omitempty"`
}

func (c *Client) ExternalIdentifiers(orcid string) (*ExternalIdentifiers, *http.Response, error) {
	data := &ExternalIdentifiers{}
	path := fmt.Sprintf("%s/external-identifiers", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) ExternalIdentifier(orcid string, putCode int) (*ExternalIdentifier, *http.Response, error) {
	data := &ExternalIdentifier{}
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddExternalIdentifier(orcid string, body *ExternalIdentifier) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/external-identifiers", orcid)
	return c.add(path, body)
}

func (c *MemberClient) UpdateExternalIdentifier(orcid string, body *ExternalIdentifier) (*ExternalIdentifier, *http.Response, error) {
	data := &ExternalIdentifier{}
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, body.PutCode)
	res, err := c.update(path, body, data)
	return data, res, err
}

func (c *MemberClient) DeleteExternalIdentifier(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, putCode)
	return c.delete(path)
}
