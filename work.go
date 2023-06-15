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

type PublicationDate struct {
	Year      *StringValue `json:"year,omitempty"`
	Month     *StringValue `json:"month,omitempty"`
	Day       *StringValue `json:"day,omitempty"`
	MediaType string       `json:"media-type,omitempty"`
}

type Citation struct {
	Type  string `json:"citation-type,omitempty"`
	Value string `json:"citation-value,omitempty"`
}

type Contributors struct {
	Contributor []Contributor `json:"contributor,omitempty"`
}

type Contributor struct {
	Attributes *ContributorAttributes `json:"contributor-attributes,omitempty"`
	CreditName *StringValue           `json:"credit-name,omitempty"`
	Email      *StringValue           `json:"contributor-email,omitempty"`
	ORCID      *URI                   `json:"contributor-orcid,omitempty"`
}

type ContributorAttributes struct {
	Role     string `json:"contributor-role,omitempty"`
	Sequence string `json:"contributor-sequence,omitempty"`
}

// TODO remaining fields
type WorkSummary struct {
	CreatedDate TimeValue `json:"created-date,omitempty"`
	// why is this string in the /work/summary api?
	DisplayIndex     int              `json:"display-index,string,omitempty"`
	ExternalIDs      *ExternalIDs     `json:"external-ids,omitempty"`
	LastModifiedDate *TimeValue       `json:"last-modified-date,omitempty"`
	Path             string           `json:"path,omitempty"`
	PublicationDate  *PublicationDate `json:"publication-date,omitempty"`
	PutCode          int              `json:"put-code,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	Title            *Title           `json:"title,omitempty"`
	Type             string           `json:"type,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
}

type GroupWork struct {
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty"`
	LastModifiedDate *TimeValue    `json:"last-modified-date,omitempty"`
	WorkSummary      []WorkSummary `json:"work-summary,omitempty"`
}

type Works struct {
	Group            []GroupWork `json:"group,omitempty"`
	LastModifiedDate *TimeValue  `json:"last-modified-date,omitempty"`
	Path             string      `json:"path,omitempty"`
}

type Work struct {
	Citation         *Citation        `json:"citation,omitempty"`
	Contributors     *Contributors    `json:"contributors,omitempty"`
	Country          *StringValue     `json:"country,omitempty"`
	CreatedDate      *TimeValue       `json:"created-date,omitempty"`
	ExternalIDs      *ExternalIDs     `json:"external-ids,omitempty"`
	JournalTitle     *StringValue     `json:"journal-title,omitempty"`
	LanguageCode     string           `json:"language-code,omitempty"`
	LastModifiedDate *TimeValue       `json:"last-modified-date,omitempty"`
	Path             string           `json:"path,omitempty"`
	PublicationDate  *PublicationDate `json:"publication-date,omitempty"`
	PutCode          int              `json:"put-code,omitempty"`
	ShortDescription string           `json:"short-description,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	Title            *Title           `json:"title,omitempty"`
	Type             string           `json:"type,omitempty"`
	URL              *StringValue     `json:"url,omitempty"`
	Visibility       string           `json:"visibility,omitempty"`
}

func (c *Client) Works(orcid string) (*Works, *http.Response, error) {
	data := &Works{}
	path := fmt.Sprintf("%s/works", orcid)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *Client) WorkSummary(orcid string, putCode int) (*WorkSummary, *http.Response, error) {
	data := &WorkSummary{}
	path := fmt.Sprintf("%s/work/summary/%d", orcid, putCode)
	res, err := c.get(path, data)
	return data, res, err
}

func (c *MemberClient) AddWork(orcid string, body *Work) (int, *http.Response, error) {
	path := fmt.Sprintf("%s/work", orcid)
	return c.add(path, body)
}
