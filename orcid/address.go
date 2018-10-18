package orcid

import (
	"fmt"
	"net/http"
)

type Address struct {
	Country          *StringValue `json:"country,omitempty"`
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	DisplayIndex     *int         `json:"display-index,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	PutCode          *int         `json:"put-code,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}
type Addresses struct {
	Address          []Address    `json:"address,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
}

func (c *Client) Addresses(orcid string) (*Addresses, *http.Response, error) {
	data := new(Addresses)
	path := fmt.Sprintf("%s/address", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) Address(orcid string, putCode int) (*Address, *http.Response, error) {
	data := new(Address)
	path := fmt.Sprintf("%s/address/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddAddress(orcid string, bodyData *Address) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/address", orcid)
	return c.add(path, bodyData)
}

func (c *MemberClient) UpdateAddress(orcid string, bodyData *Address) (*Address, *http.Response, error) {
	if err := putCodeError(bodyData.PutCode); err != nil {
		return nil, nil, err
	}
	data := new(Address)
	path := fmt.Sprintf("%s/address/%d", orcid, *bodyData.PutCode)
	res, err := c.update(path, bodyData, data)
	return data, res, err
}

func (c *MemberClient) DeleteAddress(orcid string, putCode int) (bool, *http.Response, error) {
	path := fmt.Sprintf("%s/address/%d", orcid, putCode)
	return c.delete(path)
}
