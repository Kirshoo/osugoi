package osugoi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
)

type CursorString string

type MultiplayerScores struct {
	CursorString CursorString `json:"cursor_string"`
	Params interface{} `json:"params"`
	Scores []Score `json:"scores"`
	Total *int `json:"total"`
	UserScore *Score `json:"user_score"`
}

type MultiplayerScoresAround struct {
	Higher MultiplayerScores `json:"higher"`
	Lower MultiplayerScores `json:"lower"`
}

type Score struct {
	Accuracy float32 `json:"accuracy"`
	BeatmapId int `json:"beatmap_id"`
	// TODO: Investigate BestID and BuildID
	BestId *int `json:"best_id"`
	BuildId *int `json:"build_id"`
	StartedAt *time.Time `json:"started_at"`
	EndedAt time.Time `json:"ended_at"`
	HasReplay bool `json:"has_replay"`
	Id int `json:"id"`
	IsPerfectCombo bool `json:"is_perfect_combo"`
	MaxCombo int `json:"max_combo"`
	Passed bool `json:"passed"`
	PP *float32 `json:"pp"`
	Rank string `json:"rank"`
	RulesetId int `json:"ruleset_id"` // May need to convert to string later
	TotalScore int `json:"total_score"`
	Type string `json:"type"`
	UserId int `json:"user_id"`
	
	// Only used for multiplayer score
	PlaylistItemId int `json:"playlist_item_id"`
	RoomId int `json:"room_id"`

	// Potentially depricated fields due to being for solo-score type
	ClassicTotalScore int `json:"classic_total_score"`
	Preserve bool `json:"preserve"`
	Processed bool `json:"processed"`
	Ranked bool `json:"ranked"`

	// Currently not documented types :/
	MaximumStatistics interface{} `json:"maximum_statistics"`
	Mods []interface{} `json:"mods"`
	Statistics interface{} `json:"statistics"`
	
	// Legacy fields
	LegacyPerfect bool `json:"legacy_perfect"`
	LegacyScoreId *int `json:"legacy_score_id"`
	LegacyTotalScore int `json:"legacy_total_score"`

	// Optional Attributes
	Beatmap *interface{} `json:"beatmap"`
	BeatmapSet *interface{} `json:"beatmapset"`
	// Not actually an integer!
	CurrentUserAttributes *interface{} `json:"current_user_attributes"`
	Match *interface{} `json:"match"` // Only for legacy match score
	Position *int `json:"position"` // Only for multiplayer score
	RankCountry *interface{} `json:"rank_country"`
	RankGlobal *interface{} `json:"rank_global"`
	// Only for multiplayer score
	ScoresAround *MultiplayerScoresAround `json:"scores_around"`
	User *interface{} `json:"user"`
	Weight *interface{} `json:"weight"`
}

type scoreResponse struct {
	Scores *[]Score `json:"scores"`
	CursorString CursorString `json:"cursor_string"`
}

func (c *Client) GetRecentScores(ruleset Ruleset) (*[]Score, CursorString, error) {
	req, err := c.newRequest(http.MethodGet, "/api/v2/scores", nil)
	if err != nil {
		return nil, "", fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("ruleset", string(ruleset))
	req.URL.RawQuery = query.Encode()

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("server request failed: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, "", fmt.Errorf("invalid response: %w", err)
	}

	var scoreStruct scoreResponse
	if err = json.Unmarshal(bodyBytes, &scoreStruct); err != nil {
		return nil, "", fmt.Errorf("unmarshal failed: %w", err)
	}

	return scoreStruct.Scores, scoreStruct.CursorString, nil
}
