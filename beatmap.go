package osugoi

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Ruleset string
const (
	Catch Ruleset = "fruits"
	Mania Ruleset = "mania"
	Standard Ruleset = "osu"
	Taiko Ruleset = "taiko"
)

type Failtimes struct {
	Exit *[]int `json:"exit"`
	Fail *[]int `json:"fail"`
}

type BeatmapOwner struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

type Beatmap struct {
	SetId int `json:"beatmapset_id"`
	// Potentially may be changed to float64
	Difficulty float32 `json:"difficulty_rating"`
	Id int `json:"id"`
	Mode Ruleset `json:"mode"`
	Status string `json:"status"`
	Length int `json:"total_length"`
	UserId int `json:"user_id"`
	Version string `json:"version"`

	// Optional Attributes
	// TODO:
	// Issue: is flexible. Can either be Beatmapset, BeatmapsetExtended or null!
	// Beatmapset *Beatmapset `json:"beatmapset"`
	Checksum *string `json:"checksum"`
	Playcount *int `json:"current_user_playcount"`
	Failtimes *Failtimes `json:"failtimes"`
	MaxCombo *int `json:"max_combo"`
	Owners *[]BeatmapOwner `json:"owners"`
}

func (c *Client) requestBeatmap(endpoint string) (*Beatmap, error) {
	req, err := c.newRequest(http.MethodGet, endpoint, nil)		
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	
	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server request failed: %w", err)
	}
	
	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}
	
	var beatmap Beatmap
	if err = json.Unmarshal(bodyBytes, &beatmap); err != nil {
		return nil, fmt.Errorf("beatmap unmarshal failed: %w", err)
	}

	return &beatmap, nil

}

// TODO: Add options for lookup (id, filename, etc.)
func (c *Client) LookupBeatmap() (*Beatmap, error) {
	beatmap, err := c.requestBeatmap("/api/v2/beatmaps/lookup")
	if err != nil {
		return beatmap, fmt.Errorf("lookup beatmap failed: %w", err)
	}

	return beatmap, err
}

// TODO: Change return type from Beatmap to ExtendedBeatmap
func (c *Client) GetBeatmap(beatmapId int) (*Beatmap, error) {
	beatmap, err := c.requestBeatmap(fmt.Sprintf("/api/v2/beatmaps/%d", beatmapId))
	if err != nil {
		return beatmap, fmt.Errorf("get beatmap failed: %w", err)
	}

	return beatmap, err
}
