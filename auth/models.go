package auth

const (
	defaultAuthURL string = "https://osu.ppy.sh/oauth/authorize"
	defaultTokenURL string = "https://osu.ppy.sh/oauth/token"
)

// Provides function for getting initial token
type Authenticator interface {
	Token() (*Token, error)
	RefreshToken(*Token) (*Token, error)
}

// Used to return a valid token every time
// either by recreating, refreshing or otherwise
//
// PS. This is quite similar to golang.com/x/oauth2
type TokenSource interface {
	Token() (*Token, error)
}

type AuthorizationCodeConfig struct {
	ClientId string
	ClientSecret string
	RedirectURI string
	Scopes []string

	AuthURLOverride string
	TokenURLOverride string
}

type ClientCredentialsConfig struct {
	ClientId string
	ClientSecret string

	// If left empty, will assume public
	Scopes []string

	TokenURLOverride string
}
