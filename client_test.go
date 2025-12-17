package supabase_test

import (
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

func TestNewClientWithProjectRef(t *testing.T) {
	options := &supabase.ClientOptions{
		ProjectRef: "testproject",
	}
	client, err := supabase.NewClient("https://custom.domain.com", API_KEY, options)
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}
	if client.Realtime == nil {
		t.Error("Realtime client should be initialized when ProjectRef is provided")
	}
}

func TestNewClientRealtimeNil(t *testing.T) {
	client, err := supabase.NewClient("http://localhost:54321", API_KEY, nil)
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}
	if client.Realtime != nil {
		t.Error("Realtime client should be nil for localhost URLs without ProjectRef")
	}
}

func TestNewClientRealtimeInitialized(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}
	if client.Realtime == nil {
		t.Error("Realtime client should be initialized for standard Supabase URLs")
	}
}
