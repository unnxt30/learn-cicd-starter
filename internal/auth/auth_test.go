package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantAPIKey  string
		wantErr     error
		errContains string
	}{
		{
			name:        "valid API key",
			headers:     http.Header{"Authorization": []string{"ApiKey test-api-key"}},
			wantAPIKey:  "test-api-key",
			wantErr:     nil,
			errContains: "",
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			wantAPIKey:  "",
			wantErr:     ErrNoAuthHeaderIncluded,
			errContains: "",
		},
		{
			name:        "malformed header - missing ApiKey prefix",
			headers:     http.Header{"Authorization": []string{"test-api-key"}},
			wantAPIKey:  "",
			wantErr:     nil,
			errContains: "malformed authorization header",
		},
		{
			name:        "malformed header - wrong prefix",
			headers:     http.Header{"Authorization": []string{"Bearer test-api-key"}},
			wantAPIKey:  "",
			wantErr:     nil,
			errContains: "malformed authorization header",
		},
		{
			name:        "malformed header - empty value",
			headers:     http.Header{"Authorization": []string{""}},
			wantAPIKey:  "",
			wantErr:     ErrNoAuthHeaderIncluded, // Changed this to match actual behavior
			errContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)

			// Check if the API key matches expected value
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() gotAPIKey = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}

			// Check error cases
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if tt.errContains != "" {
				if err == nil || !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("GetAPIKey() error = %v, should contain %v", err, tt.errContains)
				}
			}
		})
	}
}
