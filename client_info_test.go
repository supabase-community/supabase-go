package supabase

import (
	"strings"
	"testing"
)

func TestBuildXClientInfo(t *testing.T) {
	info := buildXClientInfo()

	if !strings.HasPrefix(info, "supabase-go/"+Version) {
		t.Fatalf("expected library token supabase-go/%s, got %q", Version, info)
	}

	required := []string{"platform=", "runtime=go", "runtime-version="}
	for _, fragment := range required {
		if !strings.Contains(info, fragment) {
			t.Errorf("expected %q in X-Client-Info, got %q", fragment, info)
		}
	}
}

func TestApplyXClientInfoHeaderSetsDefault(t *testing.T) {
	headers := map[string]string{}
	applyXClientInfoHeader(headers)

	got := headers[clientInfoHeader]
	if got == "" {
		t.Fatal("expected X-Client-Info to be set")
	}
	if !strings.HasPrefix(got, "supabase-go/") {
		t.Fatalf("unexpected header value: %q", got)
	}
}

func TestApplyXClientInfoHeaderPreservesUserValue(t *testing.T) {
	custom := "my-app/1.0; platform=edge"
	headers := map[string]string{clientInfoHeader: custom}
	applyXClientInfoHeader(headers)

	if headers[clientInfoHeader] != custom {
		t.Fatalf("expected user header to be preserved, got %q", headers[clientInfoHeader])
	}
}

func TestApplyXClientInfoHeaderPreservesUserValueCaseInsensitive(t *testing.T) {
	custom := "my-app/1.0"
	headers := map[string]string{"x-client-info": custom}
	applyXClientInfoHeader(headers)

	if headers["x-client-info"] != custom {
		t.Fatalf("expected user header to be preserved, got %q", headers["x-client-info"])
	}
	if _, ok := headers[clientInfoHeader]; ok {
		t.Fatal("expected canonical header key not to be added when user header exists")
	}
}
