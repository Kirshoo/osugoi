package auth

import (
	"fmt"
	"net/url"
	"net/http"
	"strings"
	"bytes"
	"encoding/json"
	"io"
	"time"
)

const (
	clientCredentialsGrantType string = "client_credentials"
)

func (c *ClientCredentialsConfig) tokenURL() string {
	if c.TokenURLOverride != "" {
		return c.TokenURLOverride
	}

	return defaultTokenURL
}

func (c *ClientCredentialsConfig) Token() (*Token, error) {
	data := url.Values{}

	data.Add("client_id", c.ClientId)
	data.Add("client_secret", c.ClientSecret)
	data.Add("grant_type", clientCredentialsGrantType)

	if len(c.Scopes) == 0 {
		c.Scopes = []string{"public"}
	}
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

func (c *ClientCredentialsConfig) RefreshToken(*Token) (*Token, error) {
	return c.Token()
} 
