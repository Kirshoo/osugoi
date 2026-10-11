package users

import "github.com/Kirshoo/osugoi/common"

type UserOptions struct {
	Mode common.Ruleset `query:"mode"`
	Limit int `query:"limit"`
	Offset int `query:"offset"`
	LegacyOnly int `query:"legacy_only"`
	Fails int `query:"include_fails"`
	Ids []int `query:"ids[]"`
	VariantStatistics bool `query:"include_variant_statistics"`
}

type UserOption func(*UserOptions)
func WithMode(mode common.Ruleset) UserOption {
	return func(opts *UserOptions) {
		opts.Mode = mode
	}
}

func WithLimit(limit int) UserOption {
	return func(opts *UserOptions) {
		opts.Limit = limit
	}
}

func WithOffset(offset int) UserOption {
	return func(opts *UserOptions) {
		opts.Offset = offset
	}
}

func WithIds(userIds []int) UserOption {
	return func(opts *UserOptions) {
		opts.Ids = userIds
	}
}

func LegacyOnly() UserOption {
	return func(opts *UserOptions) {
		opts.LegacyOnly = 1
	}
}

func WithFails() UserOption {
	return func(opts *UserOptions) {
		opts.Fails = 1
	}
}

func WithVariantStatistics() UserOption {
	return func(opts *UserOptions) {
		opts.VariantStatistics = true
	}
}

