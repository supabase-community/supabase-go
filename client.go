package supabase

import (
	"errors"

	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
	postgrest "github.com/supabase-community/postgrest-go"
	storage_go "github.com/supabase-community/storage-go"
)

const (
	REST_URL     = "/rest/v1"
	STORGAGE_URL = "/storage/v1"
	AUTH_URL     = "/auth/v1"
)

type Client struct {
	// Why is this a private field??
	rest    postgrest.Client
	Storage storage_go.Client
	Auth    gotrue.Client
	options clientOptions
}

type clientOptions struct {
	url     string
	headers map[string]string
}

// NewClient creates a new Supabase client.
// url is the Supabase URL.
// key is the Supabase API key.
// options is the Supabase client options.
func NewClient(url, key string, schema string, headers map[string]string) (*Client, error) {
	if url == "" || key == "" {
		return nil, errors.New("url and key are required")
	}

	if headers == nil {
		headers = map[string]string{}
	}

	headers["Authorization"] = "Bearer " + key
	headers["apikey"] = key

	if headers != nil {
		for k, v := range headers {
			headers[k] = v
		}
	}

	client := &Client{}
	client.options.url = url
	// map is pass by reference, so this gets updated by rest of function
	client.options.headers = headers

	if schema == "" {
		schema = "public"
	}

	// why pointer to an interface???
	// this isn't necessary in go
	// TODO: fix in other modules
	client.rest = *postgrest.NewClient(url+REST_URL, schema, headers)
	client.Storage = *storage_go.NewClient(url+STORGAGE_URL, key, headers)
	// ugly to make auth client use custom URL
	tmp := gotrue.New(url, key)
	client.Auth = tmp.WithCustomGoTrueURL(url + AUTH_URL)

	return client, nil
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
