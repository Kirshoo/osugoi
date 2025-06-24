package beatmappacks

import "github.com/Kirshoo/osugoi/common"

type BeatmapPackOptions struct {
	LegacyOnly int `query:"legacy_only"`
	PackType string `query:"type"`
	Cursor common.CursorString `query:"cursor_string"`
}
type BeatmapPackOption func(*BeatmapPackOptions)

func LegacyOnly() BeatmapPackOption {
	return func(opts *BeatmapPackOptions) {
		opts.LegacyOnly = 1
	}
}

func WithType(packType BeatmapPackTagType) BeatmapPackOption {
	return func(opts *BeatmapPackOptions) {
		opts.PackType = packType.Name
	}
}

func WithCursorString(cursor common.CursorString) BeatmapPackOption {
	return func(opts *BeatmapPackOptions) {
		opts.Cursor = cursor
	}
}
