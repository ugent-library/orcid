package orcid

import (
	"fmt"
	"net/http"
)

type Title struct {
	SubTitle        *StringValue `json:"subtitle,omitempty"`
	Title           *StringValue `json:"title,omitempty"`
	TranslatedTitle *StringValue `json:"translated-title,omitempty"`
}

// TODO remaining fields
type WorkSummary struct {
	CreatedDate *StringValue `json:"created-date,omitempty"`
	// why is this string in the /work/summary api?
	DisplayIndex     *int         `json:"display-index,string,omitempty"`
	ExternalIds      *ExternalIds `json:"external-ids,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Title            *Title       `json:"title,omitempty"`
}

type Work struct {
	ExternalIds      *ExternalIds  `json:"external-ids,omitempty"`
	LastModifiedDate *StringValue  `json:"last-modified-date,omitempty"`
	WorkSummary      []WorkSummary `json:"work-summary,omitempty"`
}

type Works struct {
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	Group            []Work       `json:"group,omitempty"`
}

func (c *Client) Works(orcid string) (*Works, *http.Response, error) {
	data := new(Works)
	path := fmt.Sprintf("%s/works", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) WorkSummary(orcid string, putCode int) (*WorkSummary, *http.Response, error) {
	data := new(WorkSummary)
	path := fmt.Sprintf("%s/work/summary/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}
