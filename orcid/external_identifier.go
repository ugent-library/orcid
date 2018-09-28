package orcid

import (
	"fmt"
	"net/http"
)

type ExternalIdentifier struct {
	CreatedDate  *IntValue `json:"created-date,omitempty"`
	DisplayIndex *int      `json:"display-index,omitempty"`
	// swagger docs mistakenly say object
	Relationship     *string      `json:"external-id-relationship,omitempty"`
	Type             *string      `json:"external-id-type,omitempty"`
	Value            *string      `json:"external-id-value,omitempty"`
	URL              *StringValue `json:"external-id-url,omitempty"`
	LastModifiedDate *IntValue    `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

type ExternalIdentifiers struct {
	ExternalIdentifier []ExternalIdentifier `json:"external-identifier,omitempty"`
	LastModifiedDate   *IntValue            `json:"last-modified-date,omitempty"`
	Path               *string              `json:"path,omitempty"`
}

func (c *Client) ExternalIdentifiers(orcid string) (*ExternalIdentifiers, *http.Response, error) {
	data := new(ExternalIdentifiers)
	path := fmt.Sprintf("%s/external-identifiers", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) ExternalIdentifier(orcid string, putCode int) (*ExternalIdentifier, *http.Response, error) {
	data := new(ExternalIdentifier)
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddExternalIdentifier(orcid string, bodyData *ExternalIdentifier) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/external-identifiers", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateExternalIdentifier(orcid string, bodyData *ExternalIdentifier) (*ExternalIdentifier, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(ExternalIdentifier)
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteExternalIdentifier(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/external-identifiers/%d", orcid, putCode)
	return c.delete(path)
}
