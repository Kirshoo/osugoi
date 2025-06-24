package auth

// May be changed to golang.org/x/oauth2
// if there are many requests to do so

import (
	"fmt"
	"net/url"
	"net/http"

	"strings"
	"bytes"
	"io"
	"encoding/json"

	"time"

	// Helper to open browser on auth page
	"github.com/pkg/browser"

	state "github.com/Kirshoo/osugoi/internal/stategen"
)

// If possible change to not hardcoded value
const (
	authCodeResponseType string = "code"
	authRefreshGrantType string = "refresh_token"
	authCodeGrantType string = "authorization_code"
)

func (c *AuthorizationCodeConfig) authURL() string {
	if (c.AuthURLOverride != "") {
		return c.AuthURLOverride
	}

	return defaultAuthURL
}

func (c *AuthorizationCodeConfig) tokenURL() string {
	if (c.TokenURLOverride != "") {
		return c.TokenURLOverride
	}

	return defaultTokenURL
}

func (c *AuthorizationCodeConfig) AuthCodeURL(opts ...AuthorizationOption) string {
	var options AuthorizationOptions
	for _, opt := range opts {
		opt(&options)
	}

	query := url.Values{}

	query.Add("client_id", c.ClientId)
	query.Add("redirect_uri", c.RedirectURI)
	query.Add("response_type", authCodeResponseType)
	if options.State != "" {
		query.Add("state", options.State)
	}
	query.Add("scope", strings.Join(c.Scopes, " "))

	return c.authURL() + "?" + query.Encode()
}

func (c *AuthorizationCodeConfig) Exchange(code string) (*Token, error) {
	data := url.Values{}

	data.Add("client_id", c.ClientId)
	data.Add("client_secret", c.ClientSecret)
	data.Add("redirect_uri", c.RedirectURI)
	data.Add("grant_type", authCodeGrantType)
	data.Add("code", code)

	req, err := http.NewRequest(http.MethodPost, c.tokenURL(),
			bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed request creation: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving response: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %w", err)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("error response: %d %s %s",
			resp.StatusCode, resp.Status, string(bodyBytes))
	}

	var token Token
	if err = json.Unmarshal(bodyBytes, &token); err != nil {
		return nil, fmt.Errorf("body unmarshal failed: %w", err)
	}

	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return &token, nil
}

func (c *AuthorizationCodeConfig) Token() (*Token, error) {
	// Channel to wait for user to
	// complete authorization
	codeCh := make(chan string)
	mux := http.NewServeMux()

	parsedURL, err := url.Parse(c.RedirectURI)
	if err != nil {
		return nil, fmt.Errorf("invalid redirect URI: %w", err)
	}

	path := parsedURL.Path
	addr := parsedURL.Host

	state, err := state.Generate()
	if err != nil {
		return nil, fmt.Errorf("state generation failed: %w", err)
	}

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		reqState := r.URL.Query().Get("state")
		if reqState != state {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			fmt.Println("State mismatch! Possible CSRF attack.")
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Authorization code missing", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Authorization successful. You may close this window.")
		codeCh <- code
	})

	server := &http.Server{Addr: addr, Handler: mux}
	go func() {
		_ = server.ListenAndServe()
	}()
	defer server.Close()

	authURL := c.AuthCodeURL(WithState(state))

	fmt.Println("Open this URL in your browser to authorize:")
	fmt.Println(authURL)
	browser.OpenURL(authURL)

	code := <-codeCh

	return c.Exchange(code)
}

func (c *AuthorizationCodeConfig) RefreshToken(t *Token) (*Token, error) {
	data := url.Values{}

	data.Add("client_id", c.ClientId)
	data.Add("client_secret", c.ClientSecret)
	data.Add("grant_type", authRefreshGrantType)
	data.Add("refresh_token", t.RefreshToken)
	data.Add("scope", strings.Join(c.Scopes, " "))

	req, err := http.NewRequest(http.MethodPost, c.tokenURL(),
			bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed request creation: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving response: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %w", err)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("error response: %d %s %s",
			resp.StatusCode, resp.Status, string(bodyBytes))
	}

	var token Token
	if err = json.Unmarshal(bodyBytes, &token); err != nil {
		return nil, fmt.Errorf("body unmarshal failed: %w", err)
	}

	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return &token, nil
}
