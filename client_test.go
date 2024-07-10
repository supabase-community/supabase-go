package supabase_test

import (
	"fmt"
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = "https://your-company.supabase.co"
	API_KEY = "your-api-key"
	REF_ID  = "your-project-ref-id"
)

func TestFrom(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	fmt.Println(string(data), err, count)
}

func TestRpc(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result := client.Rpc("hello_world", "", nil)
	fmt.Println(result)
}

func TestStorage(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result, err := client.Storage.GetBucket("bucket-id")
	fmt.Println(result, err)
}

func TestFunctions(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result, err := client.Functions.Invoke("hello_world", map[string]interface{}{"name": "world"})
	fmt.Println(result, err)
}

func TestRealtime(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{RefId: REF_ID})
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	err = client.Realtime.Connect()
	fmt.Println(err)
}
