package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Authorization", "ApiKey my-api-key")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if key != "my-api-key" {
		t.Fatalf("Expected API key 'my-api-key', but got: %s", key)
	}
}

func TestGetAPIKey_Failure(t *testing.T) {
	headers := make(http.Header)

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}