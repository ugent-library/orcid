package orcid

import (
	"fmt"
	"net/http"
)

type Address struct {
	Country          StringValue `json:"country,omitempty"`
	CreatedDate      TimeValue   `json:"created-date,omitempty"`
	DisplayIndex     int         `json:"display-index,omitempty"`
	LastModifiedDate TimeValue   `json:"last-modified-date,omitempty"`
	Path             string      `json:"path,omitempty"`
	PutCode          int         `json:"put-code,omitempty"`
	Source           *Source     `json:"source,omitempty"`
	Visibility       string      `json:"visibility,omitempty"`
}

type Addresses struct {
	Address          []Address `json:"address,omitempty"`
	LastModifiedDate TimeValue `json:"last-modified-date,omitempty"`
	Path             string    `json:"path,omitempty"`
}

func (c *Client) Addresses(orcid string) (*Addresses, *http.Response, error) {
	data := &Addresses{}
	path := fmt.Sprintf("%s/address", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) Address(orcid string, putCode int) (*Address, *http.Response, error) {
	data := &Address{}
	path := fmt.Sprintf("%s/address/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddAddress(orcid string, body *Address) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/address", orcid)
	return c.add(path, body)
}

func (c *MemberClient) UpdateAddress(orcid string, body *Address) (*Address, *http.Response, error) {
	data := &Address{}
	path := fmt.Sprintf("%s/address/%d", orcid, body.PutCode)
	res, err := c.update(path, body, data)
	return data, res, err
}

func (c *MemberClient) DeleteAddress(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/address/%d", orcid, putCode)
	return c.delete(path)
}
