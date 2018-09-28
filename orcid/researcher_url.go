package orcid

import (
	"fmt"
	"net/http"
)

type ResearcherURL struct {
	CreatedDate      *IntValue `json:"created-date,omitempty"`
	DisplayIndex     *int      `json:"display-index,omitempty"`
	LastModifiedDate *IntValue `json:"last-modified-date,omitempty"`
	Path             *string   `json:"path,omitempty"`
	PutCode          *int      `json:"put-code,omitempty"`
	Source           *Source   `json:"path,omitempty"`
	URLName          *string   `json:"url-name,omitempty"`
	// swagger docs mistakenly say object
	URL        *string `json:"url,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

type ResearcherURLs struct {
	LastModifiedDate *IntValue       `json:"last-modified-date,omitempty"`
	Path             *string         `json:"path,omitempty"`
	ResearcherURL    []ResearcherURL `json:"researcher-url,omitempty"`
}

func (c *Client) ResearcherURLs(orcid string) (*ResearcherURLs, *http.Response, error) {
	data := new(ResearcherURLs)
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) ResearcherURL(orcid string, putCode int) (*ResearcherURL, *http.Response, error) {
	data := new(ResearcherURL)
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddResearcherURL(orcid string, bodyData *ResearcherURL) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateResearcherURL(orcid string, bodyData *ResearcherURL) (*ResearcherURL, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(ResearcherURL)
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteResearcherURL(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/researcher-urls/%d", orcid, putCode)
	return c.delete(path)
}
