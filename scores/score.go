package scores

import (
	"net/http"
	"fmt"
	"context"

	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/common"
	"github.com/Kirshoo/osugoi/internal/options"
	"github.com/Kirshoo/osugoi/internal/optionquery"
)

const baseScoresAPI = "/api/v2/scores"

type ScoreService struct {
	Transport *transport.Transport
}

func assignOptions(opts []ScoreOption, options *ScoreOptions) {
	for _, opt := range opts {
		opt(options)
	}
}

type listScoresResponse struct {
	Scores []common.Score `json:"scores"`
	Cursor common.CursorString `json:"cursor_string"`
}

func (s *ScoreService) List(ctx context.Context, opts ...ScoreOption) (*[]common.Score, *common.CursorString, error) {
	endpointURL := baseScoresAPI
	allowedParameters := []string{"ruleset", "cursor_string"}

	req, err := s.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}

	var parameters ScoreOptions
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Accept", "application/json")

	var response listScoresResponse
	if err = s.Transport.Do(req, &response); err != nil {
		return nil, nil, fmt.Errorf("performing request: %w", err)
	}

	return &response.Scores, &response.Cursor, nil
}

// This is an undocumented endpoint and thus - is experimental
func (s *ScoreService) Get(ctx context.Context, scoreId string) (*common.Score, error) {
	endpointURL := fmt.Sprintf(baseScoresAPI + "/%s", scoreId)

	req, err := s.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var score common.Score
	if err = s.Transport.Do(req, &score); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &score, nil
}
