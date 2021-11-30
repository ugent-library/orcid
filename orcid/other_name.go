package orcid

import (
	"fmt"
	"net/http"
)

type OtherName struct {
	Content          string     `json:"content,omitempty"`
	CreatedDate      *TimeValue `json:"created-date,omitempty"`
	DisplayIndex     int        `json:"display-index,omitempty"`
	LastModifiedDate *TimeValue `json:"last-modified-date,omitempty"`
	Path             string     `json:"path,omitempty"`
	PutCode          int        `json:"put-code,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
}

type OtherNames struct {
	LastModifiedDate *TimeValue  `json:"last-modified-date,omitempty"`
	OtherName        []OtherName `json:"other-name,omitempty"`
	Path             string      `json:"path,omitempty"`
}

func (c *Client) OtherNames(orcid string) (*OtherNames, *http.Response, error) {
	data := &OtherNames{}
	path := fmt.Sprintf("%s/other-names", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) OtherName(orcid string, putCode int) (*OtherName, *http.Response, error) {
	data := &OtherName{}
	path := fmt.Sprintf("%s/other-names/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddOtherName(orcid string, body *OtherName) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/other-names", orcid)
	return c.add(path, body)
}

func (c *MemberClient) UpdateOtherName(orcid string, body *OtherName) (*OtherName, *http.Response, error) {
	data := &OtherName{}
	path := fmt.Sprintf("%s/other-names/%d", orcid, body.PutCode)
	res, err := c.update(path, body, data)
	return data, res, err
}

func (c *MemberClient) DeleteOtherName(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/other-names/%d", orcid, putCode)
	return c.delete(path)
}
