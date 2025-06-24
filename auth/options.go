package auth

type AuthorizationOptions struct {
	State string
}
type AuthorizationOption func(*AuthorizationOptions)

func WithState(state string) AuthorizationOption {
	return func(opts *AuthorizationOptions) {
		opts.State = state
	}
}
