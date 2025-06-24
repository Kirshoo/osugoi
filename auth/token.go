package auth

import "time"

type Token struct {
	Type string `json:"token_type"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`

	// time is provided in seconds
	ExpiresIn int `json:"expires_in"`
	ExpiresAt time.Time
}

func (t *Token) IsExpired() bool {
	// Add one minute buffer
	return time.Now().After(t.ExpiresAt.Add(-time.Minute))
}
