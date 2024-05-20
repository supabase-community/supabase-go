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
	rest    *postgrest.Client
	Storage *storage_go.Client
	// Auth is an interface. We don't need a pointer to an interface.
	Auth    gotrue.Client
	options clientOptions
}

type clientOptions struct {
	url     string
	headers map[string]string
}

type ClientOptions struct {
	Headers map[string]string
	Schema  string
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

	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			headers[k] = v
		}
	}

	client := &Client{}
	client.options.url = url
	// map is pass by reference, so this gets updated by rest of function
	client.options.headers = headers

	var schema string
	if options != nil && options.Schema != "" {
		schema = options.Schema
	} else {
		schema = "public"
	}

	client.rest = postgrest.NewClient(url+REST_URL, schema, headers)
	client.Storage = storage_go.NewClient(url+STORGAGE_URL, key, headers)
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

func (c *Client) SignInWithEmailPassword(email, password string) error {
	token, err := c.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return err
	}
	c.UpdateAuthSession(token.Session)

	return err
}

func (c *Client) SignInWithPhonePassword(phone, password string) error {
	token, err := c.Auth.SignInWithPhonePassword(phone, password)
	if err != nil {
		return err
	}
	c.UpdateAuthSession(token.Session)
	return err
}

func (c *Client) RefreshToken(refreshToken string) error {
	token, err := c.Auth.RefreshToken(refreshToken)
	if err != nil {
		return err
	}
	c.UpdateAuthSession(token.Session)
	return err
}

func (c *Client) UpdateAuthSession(session types.Session) {
	c.Auth = c.Auth.WithToken(session.AccessToken)
	c.rest.SetAuthToken(session.AccessToken)
	c.options.headers["Authorization"] = "Bearer " + session.AccessToken
	c.Storage = storage_go.NewClient(c.options.url, session.AccessToken, c.options.headers)

}
