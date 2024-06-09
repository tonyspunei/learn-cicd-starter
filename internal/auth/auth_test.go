package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		expectedAPIKey string
		expectedError  error
	}{
		{
			name:          "No Authorization Header",
			authHeader:    "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Authorization Header",
			authHeader:    "Bearer token",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name:           "Valid Authorization Header",
			authHeader:     "ApiKey abcdef12345",
			expectedAPIKey: "abcdef12345",
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			headers.Set("Authorization", tt.authHeader)

			apiKey, err := GetAPIKey(headers)

			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}

			if apiKey != tt.expectedAPIKey {
				t.Error("Expected API Key to be", tt.expectedAPIKey, "got", apiKey)
			}
		})
	}
}
