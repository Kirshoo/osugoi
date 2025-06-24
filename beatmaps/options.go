package beatmaps

import (
	"strconv"
	"github.com/Kirshoo/osugoi/common"
)

type BeatmapOptions struct {
	// can be represented as bool
	// 1 and 0 are true and false respectively
	LegacyOnly int `query:"legacy_only"`

	// Since some endpoints require Ruleset option
	// under mode and sum under ruleset, i decided
	// to put 2 here.
	// Potential improvement: set all 3 at the same time
	// when option is called
	Ruleset common.Ruleset `query:"ruleset"`
	Mode common.Ruleset `query:"mode"`
	RulesetId int `query:"ruleset_id"`

	// lookup options
	Id string `query:"id"`
	Checksum string `query:"checksum"`
	Filename string `query:"filename"`

	// list options
	Ids []int `query:"ids[]"`

	// Beatmap scores also accepts mods and type options,
	// but there is almost no documentation about it
	// TODO: Add Mods and Type options
	// Mods []string `query:"mods"`
	// Type string `query:"type"`
}
type BeatmapOption func(*BeatmapOptions)

func LegacyOnly() BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.LegacyOnly = 1
	}
}

func WithRuleset(mode common.Ruleset) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Ruleset = mode
	}
}

func WithMode(mode common.Ruleset) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Mode = mode
	}
}

func WithRulesetId(modeId int) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.RulesetId = modeId
	}
}

func WithId(beatmapId int) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Id = strconv.Itoa(beatmapId)
	}
}

func WithChecksum(checksum string) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Checksum = checksum
	}
}

func WithFilename(filename string) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Filename = filename
	}
}

func WithIds(beatmapIdList []int) BeatmapOption {
	return func(opts *BeatmapOptions) {
		opts.Ids = beatmapIdList
	}
}
