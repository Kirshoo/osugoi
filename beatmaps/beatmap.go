package beatmaps

import (
	"fmt"
	"net/http"
	"context"

	"github.com/Kirshoo/osugoi/client"
	"github.com/Kirshoo/osugoi/internal/optionquery"
	"github.com/Kirshoo/osugoi/internal/options"
)

const baseBeatmapAPI string = "/api/v2/beatmaps"

type Beatmaps struct {
	Client *client.Client
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

func (b *Beatmaps) Lookup(ctx context.Context, opts ...BeatmapOption) (*BeatmapExtended, error) {
	allowedParameters := []string{"id", "checksum", "filename"}

	if len(opts) == 0 {
		return nil, fmt.Errorf("at least one option must be provided")
	}

	endpointURL := baseBeatmapAPI + "/lookup"

	req, err := b.Client.NewRequest(ctx, http.MethodGet, 
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

	b.Client.Logger().Debug().Str("raw_query", req.URL.RawQuery).Msg("Request information")

	var beatmap BeatmapExtended
	if err = b.Client.Do(req, &beatmap); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmap, nil
}

func (b *Beatmaps) Get(ctx context.Context, beatmapId int) (*BeatmapExtended, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d", beatmapId)

	req, err := b.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	var beatmap BeatmapExtended
	if err = b.Client.Do(req, &beatmap); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmap, nil
}

type beatmapListResponse struct {
	Beatmaps []BeatmapExtended `json:"beatmaps"`
}

func (b *Beatmaps) List(ctx context.Context, opts ...BeatmapOption) (*[]BeatmapExtended, error) {
	endpointURL := baseBeatmapAPI
	allowedParameters := []string{"ids[]"}

	req, err := b.Client.NewRequest(ctx, http.MethodGet, 
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

	var response beatmapListResponse
	if err = b.Client.Do(req, &response); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &response.Beatmaps, nil
}

type attributesResponse struct {
	Attributes DifficultyAttributes `json:"attributes"`
}

func (b *Beatmaps) GetAttributes(ctx context.Context, beatmapId int, opts ...BeatmapOption) (*DifficultyAttributes, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/attributes", beatmapId)
	allowedParameters := []string{"mods", "ruleset", "ruleset_id"}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	data := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &data)

	req, err := b.Client.NewRequest(ctx, http.MethodPost, 
			endpointURL, data.Encode())
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var response attributesResponse
	if err = b.Client.Do(req, &response); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &response.Attributes, nil
}

func (b *Beatmaps) GetScores(ctx context.Context, beatmapId int, opts ...BeatmapOption) (*BeatmapScores, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores", beatmapId)
	allowedParameters := []string{"legacy_only", "mode", "mods", "type"}

	req, err := b.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
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
	if err = b.Client.Do(req, &scores); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &scores, nil
}

func (b *Beatmaps) GetUserScore(ctx context.Context, beatmapId, userId int, opts ...BeatmapOption) (*BeatmapUserScore, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores/users/%d", beatmapId, userId)
	allowedParameters := []string{"legacy_only", "mode", "mods"}

	req, err := b.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var userScore BeatmapUserScore
	if err = b.Client.Do(req, &userScore); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &userScore, nil
}

type allUserScores struct {
	Scores []interface{} `json:"scores"`
}

func (b *Beatmaps) GetAllUserScores(ctx context.Context, beatmapId, userId int, opts ...BeatmapOption) (*[]interface{}, error) {
	endpointURL := fmt.Sprintf(baseBeatmapAPI + "/%d/scores/users/%d/all", beatmapId, userId)
	allowedParameters := []string{"legacy_only", "ruleset"}

	req, err := b.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapOptions{}
	assignOptions(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	var userScores allUserScores
	if err = b.Client.Do(req, &userScores); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &userScores.Scores, nil
}
