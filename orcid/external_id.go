package orcid

type ExternalID struct {
	Relationship string       `json:"external-id-relationship,omitempty"`
	Type         string       `json:"external-id-type,omitempty"`
	Url          *StringValue `json:"external-id-url,omitempty"`
	Value        string       `json:"external-id-value,omitempty"`
}

type ExternalIDs struct {
	ExternalID []ExternalID `json:"external-id,omitempty"`
}
