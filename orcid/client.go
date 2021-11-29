package orcid

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
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

type Config struct {
	HTTPClient *http.Client

	ClientID     string
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
			ClientID:     cfg.ClientID,
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

var ErrNotFound = errors.New("not found")

func (c *Client) get(path string, data interface{}) (*http.Response, error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.do(req, data)
	if res.StatusCode == 404 {
		err = ErrNotFound
	} else if res.StatusCode != 200 {
		err = fmt.Errorf("couldn't get %s", path)
	}
	return res, err
}

//TODO check 201 code
func (c *MemberClient) add(path string, body interface{}) (int, *http.Response, error) {
	req, err := c.newRequest("POST", path, body)
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
func (c *MemberClient) update(path string, body, data interface{}) (*http.Response, error) {
	req, err := c.newRequest("PUT", path, body)
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
		buf = &bytes.Buffer{}
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

func IsORCID(id string) bool {
	r := regexp.MustCompile("^[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{4}$")
	return r.MatchString(id)
}
