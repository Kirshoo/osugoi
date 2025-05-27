package osugoi

import (
	"time"
	"fmt"
	"net/http"
	"encoding/json"
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

// TODO: Add options for search (BeamapPackType and Cursor)
func (c *Client) GetBeatmapPacks() (*[]BeatmapPack, error) {
	req, err := c.newRequest(http.MethodGet, "/api/v2/beatmaps/packs", nil)
	if err != nil {
		return nil, fmt.Errorf("create pack request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server pack request failed: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")

	var packs beatmapPacksResponse
	if err = json.Unmarshal(bodyBytes, &packs); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &packs.BeatmapPacks, nil
}

// TODO: add legacy_only optional parameter
func (c *Client) GetBeatmapPack(pack string) (*BeatmapPack, error) {
	req, err := c.newRequest(http.MethodGet, 
		fmt.Sprintf("/api/v2/beatmaps/packs/%s", pack), nil)
	if err != nil {
		return nil, fmt.Errorf("create pack request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server pack request failed: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")

	var beatmapPack BeatmapPack
	if err = json.Unmarshal(bodyBytes, &beatmapPack); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &beatmapPack, nil
}
