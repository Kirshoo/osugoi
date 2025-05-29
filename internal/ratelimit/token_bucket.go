package ratelimit

import (
	"time"
	"context"
)

type TokenBucket struct {
	tokens chan struct{}
}

func NewTokenBucket(ratePerSecond, burst int) *TokenBucket {
	tokenBucket := &TokenBucket{
		tokens: make(chan struct{}, burst),
	}

	for i := 0; i < burst; i++ {
		tokenBucket.tokens <- struct{}{}
	}

	go func(){
		ticker := time.NewTicker(time.Second / time.Duration(ratePerSecond))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokenBucket.tokens <- struct{}{}: // Insert new token into a bucket
			default:
			}
		}
	}()

	return tokenBucket
}

func (tb *TokenBucket) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tb.tokens:
		return nil
	}	
}
