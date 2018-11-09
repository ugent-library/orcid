package orcid

import (
	"fmt"
	"net/http"
)

type ResearcherUrl struct {
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	DisplayIndex     *int         `json:"display-index,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"path,omitempty"`
	UrlName          *string      `json:"url-name,omitempty"`
	Url              *StringValue `json:"url,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

type ResearcherUrls struct {
	LastModifiedDate *StringValue    `json:"last-modified-date,omitempty"`
	Path             *string         `json:"path,omitempty"`
	ResearcherUrl    []ResearcherUrl `json:"researcher-url,omitempty"`
}

func (c *Client) ResearcherUrls(orcid string) (*ResearcherUrls, *http.Response, error) {
	data := new(ResearcherUrls)
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) ResearcherUrl(orcid string, putCode int) (*ResearcherUrl, *http.Response, error) {
	data := new(ResearcherUrl)
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddResearcherUrl(orcid string, bodyData *ResearcherUrl) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateResearcherUrl(orcid string, bodyData *ResearcherUrl) (*ResearcherUrl, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(ResearcherUrl)
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteResearcherUrl(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	return c.delete(path)
}
