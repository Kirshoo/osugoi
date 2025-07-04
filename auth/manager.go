package auth

import (
	"fmt"
)

type TokenManager struct {
	token *Token
	config Authenticator
}

func NewTokenManager(token *Token, cfg Authenticator) *TokenManager {
	return &TokenManager{
		token: token,
		config: cfg,
	}
}

func NewTokenManagerWithAuthorization(cfg Authenticator) (*TokenManager, error) {
	token, err := cfg.Token()
	if err != nil {
		return nil, fmt.Errorf("error requesting initial token: %w", err)
	}

	return NewTokenManager(token, cfg), nil
}

func (tm *TokenManager) Token() (*Token, error) {
	if tm.token == nil {
		return nil, fmt.Errorf("authorization token is abscent or revoked")
	}

	if !tm.token.IsExpired() {
		return tm.token, nil
	}

	newToken, err := tm.config.RefreshToken(tm.token)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	tm.token = newToken
	return tm.token, nil
}

func (tm *TokenManager) RemoveToken() {
	tm.token = nil
}
