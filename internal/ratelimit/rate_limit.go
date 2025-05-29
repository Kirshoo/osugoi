package ratelimit

import (
	"fmt"
	"context"
	"net/http"
)

type RateLimiter interface {
	Acquire(ctx context.Context) error
}

type RateLimitedTransporter struct {
	Transport http.RoundTripper
	RateLimiter RateLimiter
}

func (r *RateLimitedTransporter) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := r.RateLimiter.Acquire(req.Context()); err != nil {
		return nil, fmt.Errorf("rate limited: %w", err)
	}

	transport := r.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	return transport.RoundTrip(req)
}
