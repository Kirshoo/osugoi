package osugoi

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"net/url"
	"io"
)

type clientCredentials struct {
	AccessToken string `json:"access_token"`
	ExpiersIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type Client struct {
	baseURL string
	httpAccess *http.Client
	accessToken *string
	tokenType *string
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpAccess: &http.Client{},
		accessToken: nil,
		tokenType: nil,
	}
}

// Only for Client Credential Authentications for now
func (c *Client) Authenticate(clientId, clientSecret string) error {
	data := url.Values{}
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")
	data.Add("scope", "public")

	req, err := http.NewRequest(http.MethodPost, c.baseURL + "/oauth/token",
		bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("create http request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return fmt.Errorf("server token request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("response read failed: %w", err)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errorMessage ErrorMessage
		if err = json.Unmarshal(bodyBytes, &errorMessage); err != nil {
			return fmt.Errorf("unmarshal failed: %w", err)
		}

		return &HttpRequestError{
			StatusCode: resp.StatusCode,
			Message: &errorMessage,
		}
	}

	var credentials clientCredentials
	if err = json.Unmarshal(bodyBytes, &credentials); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	c.accessToken = &credentials.AccessToken
	c.tokenType = &credentials.TokenType
	
	return nil
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL + path, body)
	if err != nil {
		return nil, fmt.Errorf("new request create failed: %w", err)
	}

	if c.accessToken != nil && c.tokenType != nil {
		req.Header.Set("Authorization", fmt.Sprintf(
			"%s %s", *c.tokenType, *c.accessToken))
	}

	return req, nil
}
