package auth_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestWrongHeader(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authentication", "123")
	_, err := auth.GetAPIKey(headers)
	if err == nil {
		log.Fatalf("Expected crash due to missing Authorization header")
	}
	if err != auth.ErrNoAuthHeaderIncluded {
		t.Fatalf("Expected %v, got: %v\n", auth.ErrNoAuthHeaderIncluded, err)
	}
}
