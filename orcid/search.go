package orcid

import (
	"fmt"
	"net/http"
	"net/url"
)

type SearchResults struct {
	NumFound int `json:"num-found,omitempty"`
	Result   []struct {
		OrcidIdentifier URI `json:"orcid-identifier,omitempty"`
	} `json:"result,omitempty"`
}

func (c *Client) Search(params url.Values) (*SearchResults, *http.Response, error) {
	path := fmt.Sprintf("search?%s", params.Encode())
	data := &SearchResults{}
	res, err := c.get(path, data)
	return data, res, err
}
