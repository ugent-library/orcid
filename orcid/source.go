package orcid

type Source struct {
	ClientID *URI         `json:"source-client-id,omitempty"`
	Name     *StringValue `json:"source-name,omitempty"`
	ORCID    *URI         `json:"source-orcid,omitempty"`
}
