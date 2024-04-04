package supabase

import (
	"errors"
	"net/url"
	"strings"

	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
	postgrest "github.com/supabase-community/postgrest-go"
	storage_go "github.com/supabase-community/storage-go"
)

const (
	REST_URL     = "/rest/v1"
	STORGAGE_URL = "/storage/v1"
)

type Client struct {
	rest    postgrest.Client
	Storage storage_go.Client
	Auth    gotrue.Client
	options clientOptions
}

type clientOptions struct {
	url     string
	headers map[string]string
}

type RestOptions struct {
	Schema string
}

type ClientOptions struct {
	Headers map[string]string
	Db      *RestOptions
}

// NewClient creates a new Supabase client.
// url is the Supabase URL.
// key is the Supabase API key.
// options is the Supabase client options.
func NewClient(url, key string, options *ClientOptions) (*Client, error) {
	if url == "" || key == "" {
		return nil, errors.New("url and key are required")
	}

	headers := map[string]string{
		"Authorization": "Bearer " + key,
		"apikey":        key,
	}

	client := &Client{}
	client.options.url = url
	// map is pass by reference, so this gets updated by rest of function
	client.options.headers = headers

	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			headers[k] = v
		}
	}

	var schema string
	if options != nil && options.Db != nil && options.Db.Schema != "" {
		schema = options.Db.Schema
	} else {
		schema = "public"
	}
	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			headers[k] = v
		}
	}

	client.rest = *postgrest.NewClient(url+REST_URL, schema, headers)
	client.Storage = *storage_go.NewClient(url+STORGAGE_URL, key, headers)

	// need reference not struct
	client.Auth = gotrue.New(getProjectReference(url), key)
	// client.Auth = tmpClient

	return client, nil
}

func getProjectReference(projectURL string) string {

	u, err := url.Parse(projectURL)
	if err != nil {
		return ""
	}
	host := u.Hostname()
	fields := strings.Split(host, ".")
	if len(fields) != 3 {
		return ""
	}

	return fields[0]

}

// Wrap postgrest From method
// From returns a QueryBuilder for the specified table.
func (c *Client) From(table string) *postgrest.QueryBuilder {
	return c.rest.From(table)
}

// Wrap postgrest Rpc method
// Rpc returns a string for the specified function.
func (c *Client) Rpc(name, count string, rpcBody interface{}) string {
	return c.rest.Rpc(name, count, rpcBody)
}

func (c *Client) SignInWithEmailPassword(email, password string) {
	token, err := c.Auth.SignInWithEmailPassword(email, password)
	c.UpdateAuthSession(token.Session, err)

}

func (c *Client) UpdateAuthSession(session types.Session, err error) {
	c.Auth = c.Auth.WithToken(session.AccessToken)
	c.rest.SetAuthToken(session.AccessToken)
	c.options.headers["Authorization"] = "Bearer " + session.AccessToken
	c.Storage = *storage_go.NewClient(c.options.url, session.AccessToken, c.options.headers)
}
