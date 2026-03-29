package testutil_test

import (
	"net/http"
	"testing"

	"github.com/supabase-community/supabase-go/internal/testutil"
)

func TestNewServerRoutesAuth(t *testing.T) {
	var gotPath string
	srv := testutil.NewServer(t, testutil.Handlers{
		Auth: func(w http.ResponseWriter, r *http.Request) {
			gotPath = r.URL.Path
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"550e8400-e29b-41d4-a716-446655440000","email":"a@b.c"}`))
		},
	})

	client := testutil.NewTestClient(t, srv, testutil.TestAPIKey)
	_, err := client.Auth.GetUser()
	if err != nil {
		t.Fatalf("GetUser: %v", err)
	}
	if gotPath != "/auth/v1/user" {
		t.Errorf("path = %q, want /auth/v1/user", gotPath)
	}
}

func TestNewServerRoutesRest(t *testing.T) {
	srv := testutil.NewServer(t, testutil.Handlers{
		Rest: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":"1"}]`))
		},
	})

	client := testutil.NewTestClient(t, srv, testutil.TestAPIKey)
	data, _, err := client.From("users").Select("*", "", false).Execute()
	if err != nil {
		t.Fatalf("Select: %v", err)
	}
	if string(data) != `[{"id":"1"}]` {
		t.Errorf("data = %q, want %q", data, `[{"id":"1"}]`)
	}
}

func TestNewServerSendsHeaders(t *testing.T) {
	srv := testutil.NewServer(t, testutil.Handlers{
		Rest: func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("apikey") != testutil.TestAPIKey {
				t.Errorf("apikey = %q, want %s", r.Header.Get("apikey"), testutil.TestAPIKey)
			}
			if r.Header.Get("Authorization") != "Bearer "+testutil.TestAPIKey {
				t.Errorf("auth = %q, want Bearer %s", r.Header.Get("Authorization"), testutil.TestAPIKey)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[]`))
		},
	})

	client := testutil.NewTestClient(t, srv, testutil.TestAPIKey)
	_, _, err := client.From("x").Select("*", "", false).Execute()
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}
}

func TestNewServerRoutesStorage(t *testing.T) {
	var gotPath string
	srv := testutil.NewServer(t, testutil.Handlers{
		Storage: func(w http.ResponseWriter, r *http.Request) {
			gotPath = r.URL.Path
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"Id":"test-bucket","Name":"test","Public":false}`))
		},
	})

	client := testutil.NewTestClient(t, srv, testutil.TestAPIKey)
	_, err := client.Storage.GetBucket("test-bucket")
	if err != nil {
		t.Fatalf("GetBucket: %v", err)
	}
	if gotPath != "/storage/v1/bucket/test-bucket" {
		t.Errorf("path = %q, want /storage/v1/bucket/test-bucket", gotPath)
	}
}

func TestNewServerRoutesFunctions(t *testing.T) {
	var gotPath string
	srv := testutil.NewServer(t, testutil.Handlers{
		Functions: func(w http.ResponseWriter, r *http.Request) {
			gotPath = r.URL.Path
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"hello"}`))
		},
	})

	client := testutil.NewTestClient(t, srv, testutil.TestAPIKey)
	_, err := client.Functions.Invoke("hello", map[string]interface{}{"name": "world"})
	if err != nil {
		t.Fatalf("Invoke: %v", err)
	}
	if gotPath != "/functions/v1/hello" {
		t.Errorf("path = %q, want /functions/v1/hello", gotPath)
	}
}
