package transport

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

type Transport struct {
	baseURL string
	tokenSource auth.TokenSource

	network *http.Client
	logger zerolog.Logger
}

func New(baseURL string, tokenSrc auth.TokenSource, opts ...TransportConfig) *Transport {
	var options TransportConfigs
	for _, opt := range opts {
		opt(&options)
	}

	t := &Transport{baseURL: baseURL, tokenSource: tokenSrc}

	httpClient := options.httpClient
	if httpClient == nil {
		httpClient = DefaultHttpClient
	}
	t.network = httpClient

	logger := options.logger
	if logger == nil {
		logger = &DefaultLogger
	}
	t.logger = *logger

	return t
}

func (t *Transport) Logger() *zerolog.Logger {
	return &t.logger
}

func (t *Transport) NewRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	url := t.baseURL + path

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

	token, err := t.tokenSource.Token();
	if err != nil {
		return nil, fmt.Errorf("request token: %w", err)
	}

	if token != nil {
		req.Header.Set("Authorization", token.Type + " " + token.AccessToken)
	}

	return req, nil
}

func (t *Transport) Do(req *http.Request, v any) error {
	resp, err := t.network.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading error body: %w", err)
		}

		t.logger.Warn().
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

// Not sure if context should be provided here, maybe create an option for it?
// For now, will use exclusively context.Background()
func (t *Transport) RevokeToken() error {
	endpointURL := "/api/v2/oauth/tokens/current"

	req, err := t.NewRequest(context.Background(), http.MethodDelete, endpointURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	// Assume no response body
	if err = t.Do(req, nil); err != nil {
		return fmt.Errorf("performing request: %w", err)
	}

	t.tokenSource.RemoveToken()
	return nil
}
