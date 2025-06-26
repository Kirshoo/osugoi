package scores

import "github.com/Kirshoo/osugoi/common"

type ScoreOptions struct {
	Ruleset common.Ruleset `query:"ruleset"`
	Cursor common.CursorString `query:"cursor_string"`
}
type ScoreOption func(*ScoreOptions)

func WithRuleset(mode common.Ruleset) ScoreOption {
	return func(opts *ScoreOptions) {
		opts.Ruleset = mode
	}
}

func WithCursorString(cursor common.CursorString) ScoreOption {
	return func(opts *ScoreOptions) {
		opts.Cursor = cursor
	}
}
