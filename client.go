package osugoi

import (
	"net/http"
	"encoding/json"
	"bytes"
	"os"
	"fmt"
	"net/url"
	"io"
)

const (
	tokenEndpoint string = "https://osu.ppy.sh/oauth/token"
)

type clientCredentials struct {
	AccessToken string `json:"access_token"`
	ExpiersIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type Client struct {
	httpAccess *http.Client
	accessToken string
	tokenType string
}

func NewClient() (*Client, error) {
	clientId, ok := os.LookupEnv("CLIENT_ID")
	if !ok {
		return nil, fmt.Errorf("cannot locate CLIENT_ID environment variable")
	}

	clientSecret, ok := os.LookupEnv("CLIENT_SECRET")
	if !ok {
		return nil, fmt.Errorf("cannot locate CLIENT_SECRET environment variable")
	}

	data := url.Values{}
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")
	data.Add("scope", "public")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, tokenEndpoint, 
		bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("create token access request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server token request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failure: %w", err)
	}

	var credentials clientCredentials
	if err = json.Unmarshal(bodyBytes, &credentials); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &Client{
		httpAccess: client,
		accessToken: credentials.AccessToken,
		tokenType: credentials.TokenType,
	}, nil
}
