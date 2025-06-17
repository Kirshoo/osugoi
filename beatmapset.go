package osugoi

import (
	"net/url"
	"reflect"
	"strings"

	"fmt"
	"net/http"
	"encoding/json"
)

type Covers struct {
	Cover string `json:"cover"`
	Cover2x string `json:"cover@2x"`
	Card string `json:"card"`
	Card2x string `json:"card@2x"`
	List string `json:"list"`
	List2x string `json:"list@2x"`
	SlimCover string `json:"slimcover"`
	SlimCover2x string `json:"slimcover@2x"`
}

type Nomination struct {
	BeatmapsetId int `json:"beatmapset_id"`
	Rulesets []Ruleset `json:"rulesets"`
	Reset bool `json:"reset"`
	UserId int `json:"user_id"`
}

type Beatmapset struct {
	Artist string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers Covers `json:"covers"`
	Creator string `json:"creator"`
	FavouriteCount int `json:"favourite_count"`
	Id int `json:"id"`
	Explicit bool `json:"nsfw"`
	Offset int `json:"offset"`
	PlayCount int `json:"play_count"`
	PreviewUrl string `json:"preview_url"`
	Source string `json:"source"`
	Status string `json:"status"`
	Spotlight bool `json:"spotlight"`
	Title string `json:"title"`
	TitleUnicode string `json:"title_unicode"`
	UserId int `json:"user_id"`
	Video bool `json:"video"`

	// Optional Fields

	// TODO: Beatmaps deserialization to a Beatmap or BeatmapExtended
	// Beatmaps *[](Beatmap || BeatmapExtended) `json:"beatmaps"`

	// TODO: Change all interface{} to an actual type
	Converts *interface{} `json:"converts"`
	CurrentNominations *[]Nomination `json:"current_nominations"`
	CurrentUserAttributes *interface{} `json:"current_user_attributes"`
	Description *interface{} `json:"description"` // Not mentioned in docs, but maybe a string?
	Discussions *interface{} `json:"discussions"`
	Events *interface{} `json:"events"`
	Genre *interface{} `json:"genre"`
	HasFavourited *bool `json:"has_favourited"`
	Language *interface{} `json:"language"`
	Nominations *interface{} `json:"nominations"` // Maybe a []Nomination type
	PackTags *[]string `json:"pack_tags"`
	Ratings *interface{} `json:"ratings"`
	RecentFavourites *interface{} `json:"recent_favourites"`
	RelatedUsers *interface{} `json:"related_users"`
	User *interface{} `json:"user"`
	TrackId *int `json:"track_id"`
}

type BeatmapsetExtended struct {
	Beatmapset
}

type SearchBeatmapsetOptions struct {
	Cursor CursorString
	limit int
}

type SearchBeatmapsetOption func(*SearchBeatmapsetOptions)
func WithCursorString(cursor CursorString) SearchBeatmapsetOption {
	return func(options *SearchBeatmapsetOptions) {
		options.Cursor = cursor
	}
}

func searchBeatmapsetOptionsToQuery(opts SearchBeatmapsetOptions) url.Values {
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

type beatmapsetSearchResponse struct {
	Beatmapsets []Beatmapset `json:"beatmapsets"`
	Cursor CursorString `json:"cursor_string"`
}

func (c *Client) SearchBeatmapset(opts ...SearchBeatmapsetOption) (*[]Beatmapset, CursorString, error) {
	req, err := c.newRequest(http.MethodGet, "/api/v2/beatmapsets/search", nil)
	if err != nil {
		return nil, "", fmt.Errorf("invalid beatmapset search request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if len(opts) > 0 {
		options := SearchBeatmapsetOptions{
			limit: 5,
		}

		for _, opt := range opts {
			opt(&options)
		}

		req.URL.RawQuery = searchBeatmapsetOptionsToQuery(options).Encode()

		c.logger.Debug().
			Str("options", fmt.Sprintf("%+v", options)).
			Str("raw_query", req.URL.RawQuery).
			Msg("Optional query")
	}

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("beatmapset search request failed: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, "", fmt.Errorf("invalid response: %w", err)
	}

	c.logger.Trace().Str("raw_response", string(bodyBytes)[:2000]).Msg("Received response")

	var response beatmapsetSearchResponse
	if err = json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, "", fmt.Errorf("beatmapset unmarshal failed: %w", err)
	}

	return &response.Beatmapsets, response.Cursor, nil
}
