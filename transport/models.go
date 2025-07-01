package transport

import (
	"os"
	"net/http"
	"github.com/rs/zerolog"

	"github.com/Kirshoo/osugoi/internal/ratelimit"
)

var (
	DefaultRateLimiter = ratelimit.NewTokenBucket(1, 2)

	DefaultHttpClient = &http.Client{
		Transport: &ratelimit.RateLimitedTransporter{
			RateLimiter: DefaultRateLimiter,
		},
	}

	DefaultLogger = zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()
)
