package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test implementation
	inputHeaders := []struct {
		name   string
		header string
	}{
		{"Valid API Key", "ApiKey valid_api_key"},
		{"Missing Authorization Header", ""},
		{"Malformed Authorization Header", "InvalidHeader"},
	}

	for _, tc := range inputHeaders {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(map[string][]string)
			if tc.header != "" {
				headers["Authorization"] = []string{tc.header}
			}

			apiKey, err := GetAPIKey(headers)
			switch tc.name {
			case "Valid API Key":
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if apiKey != "valid_api_key" {
					t.Errorf("expected apiKey to be 'valid_api_key', got '%s'", apiKey)
				}
			case "Missing Authorization Header":
				if err == nil {
					t.Error("expected an error, got nil")
				}
				if err != ErrNoAuthHeaderIncluded {
					t.Errorf("expected error to be ErrNoAuthHeaderIncluded, got %v", err)
				}
			case "Malformed Authorization Header":
				if err == nil {
					t.Error("expected an error, got nil")
				}
				if err.Error() != "malformed authorization header" {
					t.Errorf("expected error message to be 'malformed authorization header', got '%s'", err.Error())
				}
			}
		})
	}
}
