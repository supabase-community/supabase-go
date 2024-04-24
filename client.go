package supabase

import (
	"errors"

	"github.com/supabase-community/postgrest-go"
	storage "github.com/supabase-community/storage-go"
)

const (
	REST_URL      = "/rest/v1"
	STORAGE_URL   = "/storage/v1"
	DefaultSchema = "public"
)

type Client struct {
	Rest    *postgrest.Client
	Storage *storage.Client
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

	if options != nil && options.Headers != nil {
		for k, v := range options.Headers {
			headers[k] = v
		}
	}

	var schema string
	if options != nil && options.Db != nil {
		schema = options.Db.Schema
	}
	if schema == "" {
		schema = DefaultSchema
	}

	client := &Client{
		Rest:    postgrest.NewClient(url+REST_URL, schema, headers),
		Storage: storage.NewClient(url+STORAGE_URL, key, headers),
	}

	return client, nil
}

// Wrap postgrest From method
// From returns a QueryBuilder for the specified table.
func (c *Client) From(table string) *postgrest.QueryBuilder {
	return c.Rest.From(table)
}

// Wrap postgrest Rpc method
// Rpc returns a string for the specified function.
func (c *Client) Rpc(name, count string, rpcBody interface{}) string {
	return c.Rest.Rpc(name, count, rpcBody)
}
