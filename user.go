package osugoi

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"

	"github.com/Kirshoo/osugoi/internal/optionquery"
)

type User struct {
	AvatarURL string `json:"avatar_url"`
	CountryCode string `json:"country_code"`
	DefaultGroup *string `json:"default_group"`
	Id int `json:"id"`
	IsActive bool `json:"is_active"`
	IsBot bool `json:"is_bot"`
	IsDeleted bool `json:"is_deleted"`
	IsOnline bool `json:"is_online"`
	IsSupporter bool `json:"is_supporter"`
	LastVisit *time.Time `json:"last_visit"`
	PMFriendsOnly bool `json:"pm_friends_only"`
	ProfileColor string `json:"profile_colour"`
	Username string `json:"username"`

	// TODO: Add optional attributes
}

type UserExtended struct {
	User
}

type KudosuHistory struct {
	Id int `json:"id"`
	// can be 'give', 'vote.give', 'reset', 'vote.reset', 'revoke' or 'vote.revoke'
	Action string `json:"action"`
	Amount int `json:"amount"`
	Model string `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Giver *Giver `json:"giver"`
	Post Post `json:"post"`
}

type Giver struct {
	URL string `json:"url"`
	Username string `json:"username"`
}

type Post struct {
	URL *string `json:"url"`
	Title *string `json:"title"`
}

type responseKudosu struct {
	History []KudosuHistory
}

func (c *Client) GetUserKudosu(userId int) (*[]KudosuHistory, error) {
	req, err := c.newRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v2/users/%d/kudosu", userId),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpAccess.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to receive response: %w", err)
	}

	bodyBytes, err := c.getValidResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	c.logger.Trace().Str("rawResponse", string(bodyBytes)).Msg("Received body")

	var kudosu []KudosuHistory
	if err = json.Unmarshal(bodyBytes, &kudosu); err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return &kudosu, nil
}

type ScoreType string
const (
	BestScores ScoreType = "best"
	FirstScores ScoreType = "firsts"
	RecentScores ScoreType = "recent"
)

type responseUserScores struct {
	UserScores []Score
}

func (c *Client) GetUserScores(userId int, scoreType ScoreType) (*[]Score, error) {
	req, err := c.newRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v2/users/%d/scores/%s", userId, scoreType),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
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

	var scores []Score
	if err = json.Unmarshal(bodyBytes, &scores); err != nil {
		return nil, fmt.Errorf("unable to unmarshal body: %w", err)
	}

	return &scores, nil
}

type BeatmapsType string
const (
	FavouriteBeatmaps BeatmapsType = "favourite"
	GraveyardBeatmaps BeatmapsType = "graveyard"
	GuestBeatmaps BeatmapsType = "guest"
	LovedBeatmaps BeatmapsType = "loved"
	//MostPlayedBeatmaps BeatmapsType = "most_played"
	NominatedBeatmaps BeatmapsType = "nominated"
	PendingBeatmaps BeatmapsType = "pending"
	RankedBeatmaps BeatmapsType = "ranked"
)

// TODO: Allow most played beatmap types
type responseUserBeatmaps struct {
	Beatmapsets []BeatmapsetExtended
}

// TODO: Allow most played beatmap types
func (c *Client) GetUserBeatmaps(userId int, beatmaps BeatmapsType) (*[]BeatmapsetExtended, error) {
	endpointURL := fmt.Sprintf("/api/v2/users/%d/beatmapsets/%s", userId, beatmaps)
	body, err := c.doGetRaw(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}

	var beatmapsets []BeatmapsetExtended
	if err = json.Unmarshal(body, &beatmaps); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return &beatmapsets, nil
}

// Maybe split into GetUserById and GetUserByUsername?
// For now, user have to specify by prefixing user with @ sign
func (c *Client) GetUser(userId string, mode *Ruleset) (*UserExtended, error) {
	endpointURL := fmt.Sprintf("/api/v2/users/%s", userId)
	if mode != nil {
		endpointURL = endpointURL + fmt.Sprintf("/%s", *mode)
	}

	body, err := c.doGetRaw(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}

	var user UserExtended
	if err = json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return &user, nil
}

type GetUsersOptions struct {
	Ids []string `query:"ids[]"`
	IncludeVariants bool `query:"include_variant_statistics"`
}

type GetUsersOption func(*GetUsersOptions)
func WithIds(ids []string) GetUsersOption {
	return func(options *GetUsersOptions) {
		if len(ids) > 50 {
			options.Ids = ids[:50]
			return
		}

		options.Ids = ids
	}
}

func WithVariants() GetUsersOption {
	return func(options *GetUsersOptions) {
		options.IncludeVariants = true
	}
}

type responseUsers struct {
	Users []User `json:"users"`
}

func (c *Client) GetUsers(opts ...GetUsersOption) (*[]User, error) {
	endpointURL := "/api/v2/users"
	var options GetUsersOptions

	for _, opt := range opts {
		opt(&options)
	}

	query := optionquery.Convert(options)

	body, err := c.doGetRawWithQuery(endpointURL, query)
	if err != nil {
		return nil, fmt.Errorf("error requesting: %w", err)
	}

	var response responseUsers
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return &response.Users, nil
}
