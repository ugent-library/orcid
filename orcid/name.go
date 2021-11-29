package orcid

type Name struct {
	CreatedDate      TimeValue   `json:"created-date,omitempty"`
	CreditName       StringValue `json:"credit-name,omitempty"`
	FamilyName       StringValue `json:"family-name,omitempty"`
	GivenNames       StringValue `json:"given-names,omitempty"`
	LastModifiedDate TimeValue   `json:"last-modified-date,omitempty"`
	Path             string      `json:"path,omitempty"`
	Source           *Source     `json:"source,omitempty"`
	Visibility       string      `json:"visibility,omitempty"`
}
