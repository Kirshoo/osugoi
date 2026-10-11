package users

import (
	"context"
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/common"
	"github.com/Kirshoo/osugoi/internal/optionquery"
	"github.com/Kirshoo/osugoi/internal/options"
)

type UserService struct {
	Transport *transport.Transport
}

const usersBaseAPI string = "/api/v2/users"

func (u *UserService) GetSelf(ctx context.Context, opts ...UserOption) (*common.UserExtended, error) {
	endpointURL := "/api/v2/me"

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	if parameters.Mode != "" {
		endpointURL += fmt.Sprintf("/%s", parameters.Mode)
	}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var user common.UserExtended
	if err = u.Transport.Do(req, &user); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &user, nil
}

func (u *UserService) GetKudosuOf(ctx context.Context, userId int, opts ...UserOption) ([]KudosuHistory, error) {
	endpointURL := fmt.Sprintf(usersBaseAPI + "/%d/kudosu", userId)
	allowedParameters := []string{"limit", "offset"}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	query := optionquery.Convert(parameters)
	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var history []KudosuHistory
	if err = u.Transport.Do(req, &history); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return history, nil
}

type UserScoreType string
const (
	UserBestScores UserScoreType = "best"
	UserFirstScores = "firsts"
	UserRecentScores = "recent"
)

func (u *UserService) GetScoresOf(ctx context.Context, userId int, scoreType UserScoreType, opts ...UserOption) ([]common.Score, error) {
	endpointURL := fmt.Sprintf(usersBaseAPI + "/%d/scores/%s", userId, scoreType)
	allowedParameters := []string{"legacy_only", "include_fails", "mode", "limit", "offset"}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	query := optionquery.Convert(parameters)
	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var scores []common.Score
	if err = u.Transport.Do(req, &scores); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return scores, nil
}

type UserBeatmapsetType string
const (
	UserFavouriteBeatmapsets UserBeatmapsetType = "favourite"
	UserGraveyardBeatmapsets = "graveyard"
	UserGuestBeatmapsets = "guest"
	UserLovedBeatmapsets = "loved"
	UserNominatedBeatmapsets = "nominated"
	UserMostPlayedBeatmapsets = "most_played"
	UserPendingBeatmapsets = "pending"
	UserRankedBeatmapsets = "ranked"
)

func (u *UserService) getBeatmapsetsOf(ctx context.Context, userId int, beatmapsetType UserBeatmapsetType, opts ...UserOption) (*json.RawMessage, error) {
	endpointURL := fmt.Sprintf(usersBaseAPI + "/%d/beatmapsets/%s", userId, beatmapsetType)
	allowedParameters := []string{"limit", "offset"}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	query := optionquery.Convert(parameters)
	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var response json.RawMessage
	if err = u.Transport.Do(req, &response); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &response, nil
}

func (u *UserService) GetMostPlayedBeatmapsetsOf(ctx context.Context, userId int, opts ...UserOption) ([]common.BeatmapPlaycount, error) {
	body, err := u.getBeatmapsetsOf(ctx, userId, UserMostPlayedBeatmapsets, opts...)
	if err != nil {
		return nil, err
	}

	var playcounts []common.BeatmapPlaycount
	if err = json.Unmarshal(*body, &playcounts); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return playcounts, nil
}

func (u *UserService) GetCreatedBeatmapsetsOf(ctx context.Context, userId int, beatmapsetType UserBeatmapsetType, opts ...UserOption) ([]common.BeatmapsetExtended, error) {
	if beatmapsetType == UserMostPlayedBeatmapsets {
		return nil, fmt.Errorf("invalid beatmapsetType: %s. use GetMostPlayedBeatmapsetsOf() instead", beatmapsetType)
	}

	body, err := u.getBeatmapsetsOf(ctx, userId, beatmapsetType, opts...)
	if err != nil {
		return nil, err
	}

	var beatmapsets []common.BeatmapsetExtended
	if err = json.Unmarshal(*body, &beatmapsets); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return beatmapsets, nil
}

func (u *UserService) GetRecentActivityOf(ctx context.Context, userId int, opts ...UserOption) ([]common.Event, error) {
	endpointURL := fmt.Sprintf(usersBaseAPI + "/%d/recent_activity", userId)
	allowedParameters := []string{"limit", "offset"}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	query := optionquery.Convert(parameters)
	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var events []common.Event
	if err = u.Transport.Do(req, &events); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return events, nil
}

func (u *UserService) Get(ctx context.Context, userId int, opts ...UserOption) (*common.UserExtended, error) {
	endpointURL := fmt.Sprintf(usersBaseAPI + "/%d", userId)

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	if parameters.Mode != "" {
		endpointURL += fmt.Sprintf("/%s", parameters.Mode)
	}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var user common.UserExtended
	if err = u.Transport.Do(req, &user); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &user, nil
}

func (u *UserService) List(ctx context.Context, opts ...UserOption) ([]common.User, error) {
	endpointURL := usersBaseAPI
	allowedParameters := []string{"ids[]", "include_variant_statistics"}

	req, err := u.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var parameters UserOptions
	for _, opt := range opts {
		opt(&parameters)
	}

	query := optionquery.Convert(parameters)
	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var users []common.User
	if err = u.Transport.Do(req, &users); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return users, nil
}
