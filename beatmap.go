package osugoi

import (
	// For OptsToQuery helper
	"reflect"
	"net/url"
	"strings"
	"strconv"

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

type BeatmapExtended struct {
	Beatmap
}

type BeatmapPlaycount struct {
	BeatmapId int `json:"beatmap_id"`
	Beatmap *Beatmap `json:"beatmap"`
	Beatmapset *Beatmapset `json:"beatmapset"`
	Count int `json:"count"`
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

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")
	
	var beatmap Beatmap
	if err = json.Unmarshal(bodyBytes, &beatmap); err != nil {
		return nil, fmt.Errorf("beatmap unmarshal failed: %w", err)
	}

	return &beatmap, nil
}

type LookupBeatmapOptions struct {
	Id string
	Filename string
	Checksum string
}

type LookupBeatmapOption func(*LookupBeatmapOptions)

func WithId(id int) LookupBeatmapOption {
	return func(options *LookupBeatmapOptions) {
		options.Id = strconv.Itoa(id)
	}
}

func WithFilename(filename string) LookupBeatmapOption {
	return func(options *LookupBeatmapOptions) {
		options.Filename = filename
	}
}

func WithChecksum(checksum string) LookupBeatmapOption {
	return func(options *LookupBeatmapOptions) {
		options.Checksum = checksum
	}
}

// Helper to convert to valid raw query
func LookupOptionsToQuery(opts LookupBeatmapOptions) url.Values {
	v := reflect.ValueOf(opts)
	t := reflect.TypeOf(opts)

	values := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		key := strings.ToLower(fieldType.Name)

		if fieldValue.IsZero() {
			continue
		}

		values.Add(key, fieldValue.String())
	}

	return values
}

func (c *Client) LookupBeatmap(opts ...LookupBeatmapOption) (*Beatmap, error) {
	if len(opts) == 0 {
		return nil, fmt.Errorf("at least one option must be provided")
	}

	options := LookupBeatmapOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	req, err := c.newRequest(http.MethodGet, "/api/v2/beatmaps/lookup", nil)		
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	req.URL.RawQuery = LookupOptionsToQuery(options).Encode()
	
	c.logger.Debug().
		Str("options", fmt.Sprintf("%+v", options)).
		Str("raw_query", req.URL.RawQuery).
		Msg("Provided lookup options")
	
	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server request failed: %w", err)
	}
	
	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")
	
	var beatmap Beatmap
	if err = json.Unmarshal(bodyBytes, &beatmap); err != nil {
		return nil, fmt.Errorf("beatmap unmarshal failed: %w", err)
	}

	return &beatmap, nil
}

// TODO: Change return type from Beatmap to ExtendedBeatmap
func (c *Client) GetBeatmap(beatmapId int) (*Beatmap, error) {
	beatmap, err := c.requestBeatmap(fmt.Sprintf("/api/v2/beatmaps/%d", beatmapId))
	if err != nil {
		return beatmap, fmt.Errorf("get beatmap failed: %w", err)
	}

	return beatmap, err
}

type BeatmapUserScore struct {
	Position int `json:"position"`
	Score Score `json:"score"`
}

func (c *Client) GetUserBeatmapScore(beatmapId, userId int) (*BeatmapUserScore, error) {
	req, err := c.newRequest(
		http.MethodGet, 
		fmt.Sprintf("/api/v2/beatmaps/%d/scores/users/%d", beatmapId, userId),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create a request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to receive response: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")
	
	var userScore BeatmapUserScore
	if err = json.Unmarshal(bodyBytes, &userScore); err != nil {
		return nil, fmt.Errorf("userScore unmarshal failed: %w", err)
	}

	return &userScore, nil
}

type allUserScores struct {
	Scores []Score `json:"scores"`
}

// TODO: Query Parameters
func (c *Client) GetUserBeatmapScores(beatmapId, userId int) (*[]Score, error) {
	req, err := c.newRequest(
		http.MethodGet, 
		fmt.Sprintf("/api/v2/beatmaps/%d/scores/users/%d/all", beatmapId, userId),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create a request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to receive response: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")
	
	var userScores allUserScores
	if err = json.Unmarshal(bodyBytes, &userScores); err != nil {
		return nil, fmt.Errorf("userScores unmarshal failed: %w", err)
	}

	return &userScores.Scores, nil
}
