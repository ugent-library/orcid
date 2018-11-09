package orcid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	ContentType = "application/vnd.orcid+json"

	TokenUrl        = "https://orcid.org/oauth/token"
	SandboxTokenUrl = "https://sandbox.orcid.org/oauth/token"

	PublicUrl        = "https://pub.orcid.org/v2.0"
	SandboxPublicUrl = "https://pub.sandbox.orcid.org/v2.0"

	MemberUrl        = "https://api.orcid.org/v2.0"
	SandboxMemberUrl = "https://api.sandbox.orcid.org/v2.0"
)

// TODO marshalling, String() method
type Visibility int

const (
	LIMITED Visibility = iota
	REGISTERED_ONLY
	PUBLIC
	PRIVATE
)

// TODO marshalling, String() method
type ExternalIdRelationship int

const (
	SELF ExternalIdRelationship = iota
	PART_OF
)

type Config struct {
	HTTPClient *http.Client

	ClientId     string
	ClientSecret string
	Scopes       []string

	Token string

	Sandbox bool
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
}

type MemberClient struct {
	*Client
}

func newClient(baseUrl string, cfg Config) *Client {
	var httpClient *http.Client

	if cfg.HTTPClient != nil {
		httpClient = cfg.HTTPClient
	} else if cfg.Token != "" {
		t := &oauth2.Token{AccessToken: cfg.Token}
		ts := oauth2.StaticTokenSource(t)
		httpClient = oauth2.NewClient(context.Background(), ts)
	} else {
		var tokenUrl string
		if cfg.Sandbox {
			tokenUrl = SandboxTokenUrl
		} else {
			tokenUrl = TokenUrl
		}
		oauthCfg := clientcredentials.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			TokenURL:     tokenUrl,
			Scopes:       cfg.Scopes,
		}

		httpClient = oauthCfg.Client(context.Background())
	}

	return &Client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

func NewClient(cfg Config) *Client {
	if cfg.Sandbox {
		return newClient(SandboxPublicUrl, cfg)
	}
	return newClient(PublicUrl, cfg)
}

func NewMemberClient(cfg Config) *MemberClient {
	if cfg.Sandbox {
		return &MemberClient{newClient(SandboxMemberUrl, cfg)}
	}
	return &MemberClient{newClient(MemberUrl, cfg)}
}

type SearchResults struct {
	NumFound int `json:"num-found,omitempty"`
	Result   []struct {
		OrcidIdentifier Uri `json:"orcid-identifier,omitempty"`
	} `json:result,omitempty"`
}

func (c *Client) Search(q string) (*SearchResults, *http.Response, error) {
	params := url.Values{"q": {q}}
	path := fmt.Sprintf("search?%s", params.Encode())
	data := new(SearchResults)
	res, err := c.get(path, data)
	return data, res, err
}

var ErrNotFound = errors.New("Not Found")

func (c *Client) get(path string, data interface{}) (*http.Response, error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.do(req, data)
	if res.StatusCode == 404 {
		err = ErrNotFound
	} else if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Couldn't get %s", path))
	}
	return res, err
}

//TODO check 201 code
func (c *MemberClient) add(path string, bodyData interface{}) (int, *http.Response, error) {
	req, err := c.newRequest("POST", path, bodyData)
	if err != nil {
		return 0, nil, err
	}
	res, err := c.do(req, nil)
	if err != nil {
		return 0, res, err
	}
	loc, err := res.Location()
	if err != nil {
		return 0, res, err
	}
	r := regexp.MustCompile("([^/]+)$")
	match := r.FindString(loc.String())
	putCode, err := strconv.Atoi(match)
	return putCode, res, err
}

//TODO check 200 code
func (c *MemberClient) update(path string, bodyData, data interface{}) (*http.Response, error) {
	req, err := c.newRequest("PUT", path, bodyData)
	if err != nil {
		return nil, err
	}
	res, err := c.do(req, data)
	if err != nil {
		return res, err
	}
	return res, err
}

func (c *Client) delete(path string) (bool, *http.Response, error) {
	var ok bool
	req, err := c.newRequest("DELETE", path, nil)
	if err != nil {
		return ok, nil, err
	}
	res, err := c.do(req, nil)
	if res.StatusCode == 204 {
		ok = true
	}
	return ok, res, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	u := fmt.Sprintf("%s/%s", c.baseUrl, path)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return req, err
	}
	if body != nil {
		req.Header.Set("Content-Type", ContentType)
	}
	req.Header.Set("Accept", ContentType)

	return req, nil
}

func (c *Client) do(req *http.Request, data interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if data != nil {
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(data)
	}
	return res, err
}

// TODO don't expose this type
type StringValue string

func (s *StringValue) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Value interface{} `json:"value"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	var str string

	switch v := tmp.Value.(type) {
	case float64:
		str = strconv.FormatFloat(v, 'E', -1, 64)
	case int:
		str = strconv.FormatInt(int64(v), 10)
	case string:
		str = v
	default:
		return fmt.Errorf("invalid value for Value: %v of Type: %T", v)
	}

	*s = StringValue(str)

	return nil
}

type Uri struct {
	Host *string `json:"host,omitempty"`
	Path *string `json:"path,omitempty"`
	Uri  *string `json:"uri,omitempty"`
}

type Source struct {
	ClientId *Uri         `json:"source-client-id,omitempty"`
	Name     *StringValue `json:"source-name,omitempty"`
	Orcid    *Uri         `json:"source-orcid,omitempty"`
}

type ExternalId struct {
	Relationship *string      `json:"external-id-relationship,omitempty"`
	Type         *string      `json:"external-id-type,omitempty"`
	Url          *StringValue `json:"external-id-url,omitempty"`
	Value        *string      `json:"external-id-value,omitempty"`
}

type ExternalIds struct {
	ExternalId []ExternalId `json:"external-id,omitempty"`
}

type Name struct {
	CreatedDate      *StringValue `json:"created-date,omitempty"`
	CreditName       *StringValue `json:"credit-name,omitempty"`
	FamilyName       *StringValue `json:"family-name,omitempty"`
	GivenNames       *StringValue `json:"given-names,omitempty"`
	LastModifiedDate *StringValue `json:"last-modified-date,omitempty"`
	Path             *string      `json:"path,omitempty"`
	Source           *Source      `json:"path,omitempty"`
	Visibility       *string      `json:"visibility,omitempty"`
}

func Bool(v bool) *bool       { return &v }
func String(v string) *string { return &v }
func Int(v int) *int          { return &v }

func putCodeError(p *int) (err error) {
	if p == nil {
		err = errors.New("PutCode is required")
	}
	return
}

// helper functions
func IsOrcidId(id string) bool {
	r := regexp.MustCompile("^[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{4}$")
	return r.MatchString(id)
}
