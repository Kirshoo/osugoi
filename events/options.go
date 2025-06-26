package events

import "github.com/Kirshoo/osugoi/common"

type SortOption string
const (
	IdsAscending SortOption = "id_asc"
	IdsDescending SortOption = "id_desc"
)

type EventOptions struct {
	Sort SortOption `query:"sort"`
	Cursor common.CursorString `query:"cursor_string"`
}
type EventOption func(*EventOptions)

func WithCursor(cursor common.CursorString) EventOption {
	return func(opts *EventOptions) {
		opts.Cursor = cursor
	}
}

func WithSorting(sorting SortOption) EventOption {
	return func(opts *EventOptions) {
		opts.Sort = sorting
	}
}
