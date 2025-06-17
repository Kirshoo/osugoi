package osugoi

import (
	"net/http"
	"encoding/json"
	"bytes"
	"os"
	"fmt"
	"net/url"
	"io"

	"github.com/rs/zerolog"
	"github.com/Kirshoo/osugoi/internal/ratelimit"
)

type clientCredentials struct {
	AccessToken string `json:"access_token"`
	ExpiersIn int `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type Client struct {
	baseURL string
	httpAccess *http.Client
	accessToken string
	tokenType string

	logger zerolog.Logger
}

// Defaults logger to be on info level and stdout
func NewClient(baseURL string) *Client {
	return NewClientWithConfig(baseURL, nil)
}

type ClientConfigurations struct {
	Logger *zerolog.Logger
	HttpClient *http.Client
}

type ClientConfig func(*ClientConfigurations)

func WithLogger(logger zerolog.Logger) ClientConfig {
	return func(conf *ClientConfigurations) {
		conf.Logger = &logger
	}	
}

func WithHttpClient(httpClient *http.Client) ClientConfig {
	return func(conf *ClientConfigurations) {
		conf.HttpClient = httpClient
	}	
}

func NewClientWithConfig(baseURL string, configs ...ClientConfig) *Client {
	var configurations ClientConfigurations
	for _, config := range configs {
		config(&configurations)
	}

	if configurations.HttpClient == nil {
		limiter := ratelimit.NewTokenBucket(1, 2)

		defaultClient := &http.Client{
			Transport: &ratelimit.RateLimitedTransporter{
				RateLimiter: limiter,
			},
		}

		configurations.HttpClient = defaultClient
	}

	if configurations.Logger == nil {
		defaultLogger := zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()

		configurations.Logger = &defaultLogger
	}

	return &Client{
		baseURL: baseURL,
		httpAccess: configurations.HttpClient,
		logger: *configurations.Logger,
	}
}

// Only for Client Credential Authentications for now
func (c *Client) Authenticate(clientId, clientSecret string) error {
	data := url.Values{}
	data.Add("client_id", clientId)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")
	data.Add("scope", "public")

	req, err := c.newRequest(http.MethodPost, "/oauth/token",
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

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return fmt.Errorf("invalid response: %w", err)
	}
	
	var credentials clientCredentials
	if err = json.Unmarshal(bodyBytes, &credentials); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	c.accessToken = credentials.AccessToken
	c.tokenType = credentials.TokenType
	
	return nil
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	c.logger.Trace().Str("httpMethod", method).Str("path", path).Msg("Initiate new request")

	req, err := http.NewRequest(method, c.baseURL + path, body)
	if err != nil {
		return nil, fmt.Errorf("new request create failed: %w", err)
	}

	if c.accessToken != "" && c.tokenType != "" {
		authString := fmt.Sprintf("%s %s", c.tokenType, c.accessToken)
		req.Header.Set("Authorization", authString)
	}

	return req, nil
}

func (c *Client) getValidResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("response read failed: %w", err)
	}
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		c.logger.Warn().Int("statusCode", resp.StatusCode).Msg("API request")
		var errorMessage ErrorMessage
		if err = json.Unmarshal(bodyBytes, &errorMessage); err != nil {
			return bodyBytes, fmt.Errorf("unmarshal failed: %w", err)
		}

		return bodyBytes, &HttpRequestError{
			StatusCode: resp.StatusCode,
			Message: &errorMessage,
		}
	}
	
	return bodyBytes, nil
}

func (c *Client) [T any]doGet(url string) (*T, error) {
	return c.doGetWithQuery[T](url, url.Values{})
}

func (c *Client) [T any]doGetWithQuery(url, query url.Values) (*T, error) {
	req, err := c.newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.URL.RawQuery = query.Encode()

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to receive response: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("raw", string(bodyBytes)).Msg("Received body")

	var response T
	if err = json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal body: %w", err)
	}

	return &response, nil
}
