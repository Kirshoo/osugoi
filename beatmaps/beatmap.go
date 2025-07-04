package beatmaps

import (
	"fmt"
	"net/http"
	"context"

	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/internal/optionquery"
	"github.com/Kirshoo/osugoi/internal/options"
	"github.com/Kirshoo/osugoi/common"
)

const baseBeatmapAPI string = "/api/v2/beatmaps"

type BeatmapService struct {
	Transport *transport.Transport
}

// Returns error only when option parameter is
// passed as nil, otherwise never errors
func assignOptions(opts []BeatmapOption, option *BeatmapOptions) error {
	if option == nil {
		return fmt.Errorf("BeatmapOptions reference cannot be nil")
	}

	for _, opt := range opts {
		opt(option)
	}

	return nil
}

func (b *BeatmapService) Lookup(ctx context.Context, opts ...BeatmapOption) (*common.BeatmapExtended, error) {
	allowedParameters := []string{"id", "checksum", "filename"}

	if len(opts) == 0 {
		return nil, fmt.Errorf("at least one option must be provided")
	}

	endpointURL := baseBeatmapAPI + "/lookup"

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, 
			endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	b.Transport.Logger().Debug().Str("raw_query", req.URL.RawQuery).Msg("Request information")

	var beatmap common.BeatmapExtended
	if err = b.Transport.Do(req, &beatmap); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmap, nil
}

func (b *BeatmapService) Get(ctx context.Context, beatmapId int) (*common.BeatmapExtended, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d", beatmapId)

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	var beatmap common.BeatmapExtended
	if err = b.Transport.Do(req, &beatmap); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmap, nil
}

func (b *BeatmapService) List(ctx context.Context, opts ...BeatmapOption) (*[]common.BeatmapExtended, error) {
	endpointURL := baseBeatmapAPI
	allowedParameters := []string{"ids[]"}

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, 
			endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var response struct {
		Beatmaps []common.BeatmapExtended `json:"beatmaps"`
	}
	if err = b.Transport.Do(req, &response); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &response.Beatmaps, nil
}

func (b *BeatmapService) GetAttributes(ctx context.Context, beatmapId int, opts ...BeatmapOption) (*DifficultyAttributes, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/attributes", beatmapId)
	allowedParameters := []string{"mods", "ruleset", "ruleset_id"}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	data := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &data)

	req, err := b.Transport.NewRequest(ctx, http.MethodPost, 
			endpointURL, data.Encode())
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var attributes struct {
		Value DifficultyAttributes `json:"attributes"`
	}
	if err = b.Transport.Do(req, &attributes); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &attributes.Value, nil
}

func (b *BeatmapService) GetScores(ctx context.Context, beatmapId int, opts ...BeatmapOption) (*BeatmapScores, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores", beatmapId)
	allowedParameters := []string{"legacy_only", "mode", "mods", "type"}

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var scores BeatmapScores
	if err = b.Transport.Do(req, &scores); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &scores, nil
}

func (b *BeatmapService) GetUserScore(ctx context.Context, beatmapId, userId int, opts ...BeatmapOption) (*BeatmapUserScore, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores/users/%d", beatmapId, userId)
	allowedParameters := []string{"legacy_only", "mode", "mods"}

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var userScore BeatmapUserScore
	if err = b.Transport.Do(req, &userScore); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &userScore, nil
}

func (b *BeatmapService) GetAllUserScores(ctx context.Context, beatmapId, userId int, opts ...BeatmapOption) (*[]common.Score, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores/users/%d/all", beatmapId, userId)
	allowedParameters := []string{"legacy_only", "ruleset"}

	req, err := b.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var userScores struct {
		Scores []common.Score `json:"scores"`
	}
	if err = b.Transport.Do(req, &userScores); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &userScores.Scores, nil
}
