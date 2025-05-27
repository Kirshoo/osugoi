package osugoi

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestAuthenticateSuccess(t *testing.T) {
	const (
		testClientId string = "1111"
		testClientSecret string = "abcdefg"

		expectedToken = "test-token"
		expectedTokenType = "Bearer"
	)

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token_type":"Bearer", "expires_in":86400, "access_token":"test-token"}`))
	}))
	defer mockServer.Close()

	client := NewClient(mockServer.URL)
	err := client.Authenticate(testClientId, testClientSecret)
	if err != nil {
		t.Fatalf("Authorize: unexpected error: %v", err)
	}

	if client.tokenType != expectedTokenType {
		t.Errorf("Inavlid access token type: expected '%s', got '%s'",
			expectedTokenType, client.tokenType)
	}

	if client.accessToken != expectedToken {
		t.Errorf("Inavlid access token type: expected '%s', got '%s'",
			expectedToken, client.accessToken)
	}
}
