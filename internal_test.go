package supabase

import "testing"

func TestExtractProjectRef(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "standard supabase.co URL",
			url:      "https://abc.supabase.co",
			expected: "abc",
		},
		{
			name:     "supabase.co URL with path",
			url:      "https://myproject.supabase.co/rest/v1",
			expected: "myproject",
		},
		{
			name:     "supabase.in regional URL",
			url:      "https://abc.supabase.in",
			expected: "abc",
		},
		{
			name:     "custom domain",
			url:      "https://abc.example.com",
			expected: "abc",
		},
		{
			name:     "localhost URL",
			url:      "http://localhost:54321",
			expected: "",
		},
		{
			name:     "localhost without port",
			url:      "http://localhost",
			expected: "",
		},
		{
			name:     "IPv4 address",
			url:      "http://127.0.0.1:54321",
			expected: "",
		},
		{
			name:     "IPv6 address",
			url:      "http://[::1]:54321",
			expected: "",
		},
		{
			name:     "invalid URL",
			url:      "not-a-url",
			expected: "",
		},
		{
			name:     "empty URL",
			url:      "",
			expected: "",
		},
		{
			name:     "single label hostname",
			url:      "http://supabase",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractProjectRef(tt.url)
			if result != tt.expected {
				t.Errorf("extractProjectRef(%q) = %q, want %q", tt.url, result, tt.expected)
			}
		})
	}
}
