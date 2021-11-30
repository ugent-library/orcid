package orcid

import (
	"fmt"
	"net/http"
)

type Keyword struct {
	Content          string     `json:"content,omitempty"`
	CreatedDate      *TimeValue `json:"created-date,omitempty"`
	DisplayIndex     int        `json:"display-index,omitempty"`
	LastModifiedDate *TimeValue `json:"last-modified-date,omitempty"`
	Path             string     `json:"path,omitempty"`
	PutCode          int        `json:"put-code,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
}

type Keywords struct {
	Keyword          []Keyword  `json:"keyword,omitempty"`
	LastModifiedDate *TimeValue `json:"last-modified-date,omitempty"`
	Path             string     `json:"path,omitempty"`
}

func (c *Client) Keywords(orcid string) (*Keywords, *http.Response, error) {
	data := &Keywords{}
	path := fmt.Sprintf("%s/keywords", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddKeyword(orcid string, body *Keyword) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/keywords", orcid)
	return c.add(path, body)
}

func (c *MemberClient) UpdateKeyword(orcid string, body *Keyword) (*Keyword, *http.Response, error) {
	data := &Keyword{}
	path := fmt.Sprintf("%s/keywords/%d", orcid, body.PutCode)
	res, err := c.update(path, body, data)
	return data, res, err
}

func (c *MemberClient) DeleteKeyword(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/keywords/%d", orcid, putCode)
	return c.delete(path)
}
