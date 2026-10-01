package supabase

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	clientInfoHeader = "X-Client-Info"
	// Version is the supabase-go library version sent in X-Client-Info.
	Version = "0.0.5"
)

func buildXClientInfo() string {
	parts := []string{fmt.Sprintf("supabase-go/%s", Version)}

	if platform := runtime.GOOS; platform != "" {
		parts = append(parts, "platform="+platform)
	}

	parts = append(parts, "runtime=go")

	if v := strings.TrimPrefix(runtime.Version(), "go"); v != "" {
		parts = append(parts, "runtime-version="+v)
	}

	return strings.Join(parts, "; ")
}

func applyXClientInfoHeader(headers map[string]string) {
	if hasClientInfoHeader(headers) {
		return
	}
	headers[clientInfoHeader] = buildXClientInfo()
}

func hasClientInfoHeader(headers map[string]string) bool {
	for k := range headers {
		if strings.EqualFold(k, clientInfoHeader) {
			return true
		}
	}
	return false
}
