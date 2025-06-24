package auth

import (
	"fmt"
)

type TokenManager struct {
	token *Token
	config Authenticator
}

func NewTokenManager(cfg Authenticator) (*TokenManager, error) {
	token, err := cfg.Token()
	if err != nil {
		return nil, fmt.Errorf("error requesting initial token: %w", err)
	}

	return &TokenManager{
		token: token,
		config: cfg,
	}, nil
}

func (tm *TokenManager) Token() (*Token, error) {
	if tm.token.IsExpired() {
		return tm.token, nil
	}

	newToken, err := tm.config.RefreshToken(tm.token)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	tm.token = newToken
	return tm.token, nil
}
