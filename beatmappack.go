package osugoi

import (
	"time"
	"fmt"
	"encoding/json"

	"github.com/Kirshoo/osugoi/internal/optionquery"
)

type BeatmapPackType string
const (
	StandardType BeatmapPackType = "standard"
	FeaturedArtistType BeatmapPackType = "featured"
	TournamentType BeatmapPackType = "tournament"
	ProjectLovedType BeatmapPackType = "loved"
	SpotlightType BeatmapPackType = "chart"
	ThemeType BeatmapPackType = "theme"
	ArtistType BeatmapPackType = "artist"
)

type BeatmapPack struct {
	Author string `json:"author"`
	Date time.Time `json:"date"`
	Name string `json:"name"`
	DiffReduction bool `json:"no_diff_reduction"`
	RulesetId int `json:"ruleset_id"` // Possibly change to Ruleset type
	Tag string `json:"tag"` // Maybe implement translation to pack type
	DownloadUrl string `json:"url"`

	// Optional attributes

	// TODO: Add beatmapset type
	// BeatmapSets *[]Beatmapset `json:"beatmapsets"`
	CompletedBeatmapsetIds []int `json:"user_completion_data.beatmapset_ids"`
	IsCompleted bool `json:"user_completion_data.completed"`
}

type beatmapPacksResponse struct {
	BeatmapPacks []BeatmapPack `json:"beatmap_packs"`
}

type BeatmapPacksOptions struct {
	Type BeatmapPackType `query:"type"`
	Cursor CursorString `query:"cursor_string"`
}
type BeatmapPacksOption func(*BeatmapPacksOptions)

func WithPackCursor(cursor CursorString) BeatmapPacksOption {
	return func(options *BeatmapPacksOptions) {
		options.Cursor = cursor
	}
}

func WithPackType(packType BeatmapPackType) BeatmapPacksOption {
	return func(options *BeatmapPacksOptions) {
		options.Type = packType
	}
}

func (c *Client) GetBeatmapPacks(opts ...BeatmapPacksOption) (*[]BeatmapPack, error) {
	endpointURL := "/api/v2/beatmaps/packs"

	options := BeatmapPacksOptions{
		Type: StandardType,
	}

	for _, opt := range opts {
		opt(&options)
	}

	query := optionquery.Convert(options)

	bodyBytes, err := c.doGetRawWithQuery(endpointURL, query)
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}

	var packs beatmapPacksResponse
	if err = json.Unmarshal(bodyBytes, &packs); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return &packs.BeatmapPacks, nil
}

type BeatmapPackOptions struct {
	LegacyOnly int `query:"legacy_only"`
}

type BeatmapPackOption func(*BeatmapPackOptions)

func LegacyOnly() BeatmapPackOption {
	return func(options *BeatmapPackOptions) {
		options.LegacyOnly = 1
	}
}

func (c *Client) GetBeatmapPack(pack string, opts ...BeatmapPackOption) (*BeatmapPack, error) {
	endpointURL := fmt.Sprintf("/api/v2/beatmaps/packs/%s", pack)

	options := BeatmapPackOptions{
		LegacyOnly: 0,
	}

	for _, opt := range opts {
		opt(&options)
	}

	query := optionquery.Convert(options)

	bodyBytes, err := c.doGetRawWithQuery(endpointURL, query)
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}

	var responsePack BeatmapPack
	if err = json.Unmarshal(bodyBytes, &responsePack); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return &responsePack, nil
}
