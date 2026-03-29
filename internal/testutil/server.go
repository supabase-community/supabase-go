package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	supabase "github.com/supabase-community/supabase-go"
)

const TestAPIKey = "test-api-key"

type Handlers struct {
	Auth      http.HandlerFunc
	Rest      http.HandlerFunc
	Storage   http.HandlerFunc
	Functions http.HandlerFunc
}

func NewServer(t *testing.T, h Handlers) *httptest.Server {
	t.Helper()

	var mu sync.Mutex
	var unexpected []string

	mux := http.NewServeMux()
	if h.Auth != nil {
		mux.HandleFunc("/auth/v1/", h.Auth)
	}
	if h.Rest != nil {
		mux.HandleFunc("/rest/v1/", h.Rest)
	}
	if h.Storage != nil {
		mux.HandleFunc("/storage/v1/", h.Storage)
	}
	if h.Functions != nil {
		mux.HandleFunc("/functions/v1/", h.Functions)
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		unexpected = append(unexpected, fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		mu.Unlock()
		http.NotFound(w, r)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(func() {
		srv.Close()
		mu.Lock()
		defer mu.Unlock()
		for _, req := range unexpected {
			t.Errorf("unexpected request: %s", req)
		}
	})
	return srv
}

func NewTestClient(t *testing.T, srv *httptest.Server, key string) *supabase.Client {
	t.Helper()
	client, err := supabase.NewClient(srv.URL, key, nil)
	if err != nil {
		t.Fatalf("testutil.NewTestClient: %v", err)
	}
	return client
}
