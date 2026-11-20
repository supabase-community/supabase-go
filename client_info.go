package supabase

import (
	"fmt"
	"strings"
)

const (
	clientInfoHeader = "X-Client-Info"
	// Version is the supabase-go library version sent in X-Client-Info.
	Version = "0.0.5"
)

// buildXClientInfo returns the X-Client-Info value this library sends: the library token alone, with no structured metadata after it.
func buildXClientInfo() string {
	return fmt.Sprintf("supabase-go/%s", Version)
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
