package client

import (
	"net/http"
	"github.com/rs/zerolog"
)

type ClientConfigs struct {
	httpClient *http.Client
	logger *zerolog.Logger
}
type ClientConfig func(*ClientConfigs)

func WithHTTPClient(client *http.Client) ClientConfig {
	return func(cfgs *ClientConfigs) {
		cfgs.httpClient = client
	}
}

func WithLogger(logger *zerolog.Logger) ClientConfig {
	return func(cfgs *ClientConfigs) {
		cfgs.logger = logger
	}
}
