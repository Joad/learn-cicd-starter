package auth

import (
	"net/http"
	"testing"
)

func TestValidAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey testKey")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal("Failed to get valid key", err)
	}

	if key != "testKey" {
		t.Fatal("Wrong key extracted")
	}
}

func TestNoApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "testKey")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected an error, got none")
	}
}

func TestNoAuthHeaders(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}
