package supabase_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = "https://your-company.supabase.co"
	API_KEY = "your-api-key"
)

func TestNewClient(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	t.Logf("newClient result: %v", client)
}

func TestFrom(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	if err != nil {
		t.Errorf("cannot perform From operation: %v", err)
	}
	t.Logf("%s%v%d", data, err, count)
}

func TestRpc(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	result := client.Rpc("hello_world", "", nil)
	t.Logf("rpc result: %s", result)
}

func TestStorage(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	bucket, err := client.Storage.GetBucket("bucket-id")
	if err != nil {
		t.Errorf("cannot get bucket: %v", err)
	}
	t.Logf("getBucket result: %v", bucket)
}

func TestFunctions(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	result, err := client.Functions.Invoke("hello_world", map[string]interface{}{"name": "world"})
	if err != nil {
		t.Errorf("cannot invoke function: %v", err)
	}
	t.Logf("function invokation result: %v", result)
}

func TestGetUserSendsAuthHeader(t *testing.T) {
	var receivedAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedAuth = r.Header.Get("Authorization")
		if r.URL.Path == "/auth/v1/user" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"550e8400-e29b-41d4-a716-446655440000","email":"test@example.com"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client, err := supabase.NewClient(server.URL, "test-api-key", nil)
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	_, err = client.Auth.GetUser()
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}
	if receivedAuth != "Bearer test-api-key" {
		t.Errorf("expected Authorization header 'Bearer test-api-key', got %q", receivedAuth)
	}
}
