package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Test case: empty auth header should return an error
    t.Run("Empty authorization header", func(t *testing.T) {
        // Create an empty header
        headers := http.Header{}

        // Call GetAPIKey with the empty header
        key, err := GetAPIKey(headers)

        // Check if we got an error (we expect one)
        if err == nil {
            t.Errorf("Expected an error for empty auth header, but got nil")
        }

        // Check if the key is empty (it should be)
        if key != "" {
            t.Errorf("Expected empty key for empty auth header, but got: %s", key)
        }
    })

    // Test case: valid auth header with "ApiKey" prefix
    t.Run("Valid authorization header with ApiKey prefix", func(t *testing.T) {
        headers := http.Header{}
        headers.Add("Authorization", "ApiKey test-api-key")

        key, err := GetAPIKey(headers)

        if err != nil {
            t.Errorf("Expected no error for valid auth header, but got: %v", err)
        }

        if key != "test-api-key" {
            t.Errorf("Expected key to be 'test-api-key', but got: %s", key)
        }
    })

}
