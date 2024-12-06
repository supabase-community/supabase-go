package supabase_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = "https://your-company.supabase.co"
	API_KEY = "your-api-key"
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

func TestFromWithProxy(t *testing.T) {
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"name":"Test Country"}]`))
	}))
	defer proxy.Close()

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{Proxy: proxy.URL})
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	if err != nil {
		t.Fatalf("error executing query: %v", err)
	}

	expected := `[{"id":1,"name":"Test Country"}]`
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, string(data))
	}
	if count != 1 {
		t.Errorf("expected count 1, got %d", count)
	}
}

func TestRpcWithProxy(t *testing.T) {
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"Hello, world!"`))
	}))
	defer proxy.Close()

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{Proxy: proxy.URL})
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	result := client.Rpc("hello_world", "", nil)
	expected := `"Hello, world!"`
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestStorageWithProxy(t *testing.T) {
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"bucket-id","name":"Test Bucket"}`))
	}))
	defer proxy.Close()

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{Proxy: proxy.URL})
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	result, err := client.Storage.GetBucket("bucket-id")
	if err != nil {
		t.Fatalf("error getting bucket: %v", err)
	}

	expected := `{"id":"bucket-id","name":"Test Bucket"}`
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestFunctionsWithProxy(t *testing.T) {
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello, world!"}`))
	}))
	defer proxy.Close()

	client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{Proxy: proxy.URL})
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	result, err := client.Functions.Invoke("hello_world", map[string]interface{}{"name": "world"})
	if err != nil {
		t.Fatalf("error invoking function: %v", err)
	}

	expected := `{"message":"Hello, world!"}`
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
