package client

import (
	"net/http"
	"context"
	"io"
	"bytes"
	"strings"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"

	"github.com/Kirshoo/osugoi/auth"
)

type Client struct {
	baseURL string
	tokenSource auth.TokenSource

	network *http.Client
	logger zerolog.Logger
}

func New(baseURL string, tokenSrc auth.TokenSource, opts ...ClientConfig) *Client {
	var options ClientConfigs
	for _, opt := range opts {
		opt(&options)
	}

	c := &Client{baseURL: baseURL, tokenSource: tokenSrc}

	httpClient := options.httpClient
	if httpClient == nil {
		httpClient = DefaultHttpClient
	}
	c.network = httpClient

	logger := options.logger
	if logger == nil {
		logger = &DefaultLogger
	}
	c.logger = *logger

	return c
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	url := c.baseURL + path

	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encode body failure: %w", err)
		}
		buf = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token, err := c.tokenSource.Token(); err != nil {
		req.Header.Set("Authorization", token.Type + " " + token.AccessToken)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v any) error {
	resp, err := c.network.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading error body: %w", err)
		}

		c.logger.Warn().
			Str("status", resp.Status).
			Int("code", resp.StatusCode).
			Str("body", strings.TrimSpace(string(bodyBytes))).
			Msg("API response status")

		return fmt.Errorf("api error: %s", resp.Status)
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return nil
}
