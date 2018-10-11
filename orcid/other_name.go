package orcid

import (
	"fmt"
	"net/http"
)

type OtherName struct {
	Content          *string      `json:"content,omitempty"`
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	DisplayIndex     *int         `json:"display-index,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"path,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

type OtherNames struct {
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	OtherName        []OtherName  `json:"other-name,omitempty"`
	Path             *string      `json:"path,omitempty"`
}

func (c *Client) OtherNames(orcid string) (*OtherNames, *http.Response, error) {
	data := new(OtherNames)
	path := fmt.Sprintf("%s/other-names", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) OtherName(orcid string, putCode int) (*OtherName, *http.Response, error) {
	data := new(OtherName)
	path := fmt.Sprintf("%s/other-names/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddOtherName(orcid string, bodyData *OtherName) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/other-names", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateOtherName(orcid string, bodyData *OtherName) (*OtherName, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(OtherName)
	path := fmt.Sprintf("%s/other-names/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteOtherName(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/other-names/%d", orcid, putCode)
	return c.delete(path)
}
