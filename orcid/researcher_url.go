package orcid

import (
	"fmt"
	"net/http"
)

type ResearcherURL struct {
	CreatedDate      *TimeValue   `json:"created-date,omitempty"`
	DisplayIndex     int          `json:"display-index,omitempty"`
	LastModifiedDate *TimeValue   `json:"last-modified-date,omitempty"`
	Path             string       `json:"path,omitempty"`
	PutCode          int          `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	UrlName          string       `json:"url-name,omitempty"`
	Url              *StringValue `json:"url,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
}

type ResearcherURLs struct {
	LastModifiedDate *TimeValue      `json:"last-modified-date,omitempty"`
	Path             string          `json:"path,omitempty"`
	ResearcherURL    []ResearcherURL `json:"researcher-url,omitempty"`
}

func (c *Client) ResearcherURLs(orcid string) (*ResearcherURLs, *http.Response, error) {
	data := &ResearcherURLs{}
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) ResearcherURL(orcid string, putCode int) (*ResearcherURL, *http.Response, error) {
	data := &ResearcherURL{}
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddResearcherURL(orcid string, body *ResearcherURL) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	return c.add(path, body)
}

func (c *MemberClient) UpdateResearcherURL(orcid string, body *ResearcherURL) (*ResearcherURL, *http.Response, error) {
	data := &ResearcherURL{}
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, body.PutCode)
	res, err := c.update(path, body, data)
	return data, res, err
}

func (c *MemberClient) DeleteResearcherURL(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	return c.delete(path)
}
