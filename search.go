package orcid

import (
	"fmt"
	"net/http"
	"net/url"
)

type SearchResult struct {
	ORCIDIdentifier *URI `json:"orcid-identifier,omitempty"`
}

type SearchResults struct {
	NumFound int            `json:"num-found,omitempty"`
	Result   []SearchResult `json:"result,omitempty"`
}

func (c *Client) Search(params url.Values) (*SearchResults, *http.Response, error) {
	path := fmt.Sprintf("search?%s", params.Encode())
	data := &SearchResults{}
	res, err := c.get(path, data)
	return data, res, err
}
