package auth

import (
	"net/http"
	"regexp"
	"testing"
)

// Example regex pattern for a simple API key format
var apiKeyPattern = `^Bearer [a-zA-Z0-9\-_.]+$`

func TestGetAPIKey(t *testing.T) {
    tests := map[string]struct {
        input http.Header
        want  string
    }{
        "no auth header": {
            input: http.Header{},
            want:  ErrNoAuthHeaderIncluded.Error(),
        },
        "with auth header": {
            input: http.Header{"Authorization": {"Bearer some-api-key"}},
            want:  "some-api-key",
        },
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            // Call GetAPIKey
            got, err := GetAPIKey(tc.input)
            if err != nil && err.Error() != tc.want {
                t.Errorf("expected error %v, but got %v", tc.want, err)
            }
            
            // Use regex to validate the format of the result
            matched, _ := regexp.MatchString(apiKeyPattern, "Bearer "+got)
            if tc.want != "" && !matched {
                t.Errorf("expected API key to match pattern, but got %v", got)
            }
        })
    }
}