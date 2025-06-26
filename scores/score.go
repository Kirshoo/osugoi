package scores

import (
	"net/http"
	"fmt"
	"context"

	"github.com/Kirshoo/osugoi/client"
	"github.com/Kirshoo/osugoi/common"
	"github.com/Kirshoo/ousgoi/internal/options"
	"github.com/Kirshoo/ousgoi/internal/optionquery"
)

const baseScoresAPI = "/api/v2/scores"

type ScoreService struct {
	Client *client.Client
}

func assignOptions(opts []ScoreOption, options *ScoreOptions) {
	for _, opt := range opts {
		opts(options)
	}
}

type listScoresResponse struct {
	Scores []common.Score `json:"scores"`
	Cursor common.CursorString `json:"cursor_string"`
}

func (s *ScoreService) List(ctx context.Context, opts ...ScoreOption) (*[]common.Score, *common.CursorString, error) {
	endpointURL := baseScoresAPI
	allowedParameters := []string{"ruleset", "cursor_string"}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
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
	if err = s.Client.Do(req, &response); err != nil {
		return nil, nil, fmt.Errorf("performing request: %w", err)
	}

	return &response.Scores, &response.Cursor, nil
}
