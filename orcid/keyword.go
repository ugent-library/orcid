package orcid

import (
	"fmt"
	"net/http"
)

type Keyword struct {
	Content          *string      `json:"content,omitempty"`
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	DisplayIndex     *int         `json:"display-index,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

type Keywords struct {
	Keyword          []Keyword    `json:"keyword,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
}

func (c *Client) Keywords(orcid string) (*Keywords, *http.Response, error) {
	data := new(Keywords)
	path := fmt.Sprintf("%s/keywords", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddKeyword(orcid string, bodyData *Keyword) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/keywords", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateKeyword(orcid string, bodyData *Keyword) (*Keyword, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(Keyword)
	path := fmt.Sprintf("%s/keywords/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteKeyword(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/keywords/%d", orcid, putCode)
	return c.delete(path)
}
