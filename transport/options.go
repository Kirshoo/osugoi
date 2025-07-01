package transport

import (
	"net/http"
	"github.com/rs/zerolog"
)

type TransportConfigs struct {
	httpClient *http.Client
	logger *zerolog.Logger
}
type TransportConfig func(*TransportConfigs)

func WithHTTPClient(client *http.Client) TransportConfig {
	return func(cfgs *TransportConfigs) {
		cfgs.httpClient = client
	}
}

func WithLogger(logger *zerolog.Logger) TransportConfig {
	return func(cfgs *TransportConfigs) {
		cfgs.logger = logger
	}
}
