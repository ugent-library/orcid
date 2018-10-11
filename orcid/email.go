package orcid

import (
	"fmt"
	"net/http"
)

type Email struct {
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	Email            *string      `json:"email,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	Primary          *bool        `json:"primary,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Verified         *bool        `json:"verified,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

type Emails struct {
	Email            []Email      `json:"email,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
}

func (c *Client) Emails(orcid string) (*Emails, *http.Response, error) {
	data := new(Emails)
	path := fmt.Sprintf("%s/email", orcid)
	res, err := c.get(path, data)
	return data, res, err
}
