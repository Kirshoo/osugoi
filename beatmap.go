package osugoi

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

// Ruleset enum
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

func (c *Client) GetBeatmap(beatmapId int) (*Beatmap, error) {
	req, err := c.newRequest(http.MethodGet, 
		fmt.Sprintf("/api/v2/beatmaps/%d", beatmapId), nil)		
	if err != nil {
		return nil, fmt.Errorf("create beatmap request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	
	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("beatmap request failed: %w", err)
	}
	defer resp.Body.Close()

	// TODO: check response status code

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %w", err)
	}

	var beatmap Beatmap
	if err = json.Unmarshal(bodyBytes, &beatmap); err != nil {
		return nil, fmt.Errorf("beatmap unmarshal failed: %w", err)
	}

	return &beatmap, nil
}
