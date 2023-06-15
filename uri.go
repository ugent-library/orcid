package orcid

type URI struct {
	Host string `json:"host,omitempty"`
	Path string `json:"path,omitempty"`
	URI  string `json:"uri,omitempty"`
}
