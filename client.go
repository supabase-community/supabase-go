package supabase

import (
	"errors"

	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase/postgrest-go"
)

const (
	REST_URL     = "/rest/v1"
	STORGAGE_URL = "/storage/v1"
)

type Client struct {
	rest    *postgrest.Client
	Storage *storage_go.Client
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

	defaultHeaders := map[string]string{
		"Authorization": "Bearer " + key,
		"apikey":        key,
	}

	client := &Client{}
	schema := "public"
	if options != nil && options.Db != nil {
		if len(options.Headers) > 0 {
			for k, v := range options.Headers {
				defaultHeaders[k] = v
			}
		}
	}
	client.rest = postgrest.NewClient(url+REST_URL, schema, defaultHeaders)
	client.Storage = storage_go.NewClient(url+STORGAGE_URL, key, defaultHeaders)

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
