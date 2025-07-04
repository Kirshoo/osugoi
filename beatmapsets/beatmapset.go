package beatmapsets

import (
	"context"
	"net/http"
	"fmt"
	"net/url"

	"github.com/Kirshoo/osugoi/common"
	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/internal/optionquery"
)

const beatmapsetsBaseAPI string = "/api/v2/beatmapsets"

type BeatmapsetService struct {
	Transport *transport.Transport
}

type cursor struct {
	// When search.sort is "ranked_desc" or "ranked_asc". maybe something else too
	ApprovedDate float64 `json:"approved_date"` // time.Time maybe (it looks like unix timestamp)

	// When search.sort is "relevance_desc" or "relevance_asc". maybe somthing else too
	Score float64 `json:"_score"`
	Id int `json:"id"`
}

type search struct {
	// Method used to sort the map order
	Sort string `json:"sort"`
}

type searchResponse struct {
	Beatmapsets []common.BeatmapsetExtended `json:"beatmapsets"`

	// Probably is depricated, use CursorString intead
	Cursor cursor `json:"cursor"`

	CursorString common.CursorString `json:"cursor_string"`
	Error *any `json:"error"`

	// Only available when have identify scope on
	// authentication code grant
	RecommendedDifficulty *float64 `json:"recommended_difficulty"`
	Search search `json:"search"`
	Total int `json:"total"`
}

func (bs *BeatmapsetService) Search(ctx context.Context, subquery ...SubqueryParameter) (*searchResponse, error) {
	endpointURL := beatmapsetsBaseAPI + "/search"

	req, err := bs.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	query := url.Values{}
	query.Set("q", BuildSubquery(subquery...))

	bs.Transport.Logger().Debug().Str("raw_query", query.Encode()).Msg("Search request parameters")
	req.URL.RawQuery = query.Encode()

	var response searchResponse
	if err = bs.Transport.Do(req, &response); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &response, nil
}

func (bs *BeatmapsetService) Lookup(ctx context.Context, opts ...BeatmapsetOption) (*common.BeatmapsetExtended, error) {
	if len(opts) == 0 {
		return nil, fmt.Errorf("at least one optional query must be provied")
	}

	endpointURL := beatmapsetsBaseAPI + "/lookup"

	req, err := bs.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var params BeatmapsetOptions
	for _, opt := range opts {
		opt(&params)
	}

	query := optionquery.Convert(params)
	req.URL.RawQuery = query.Encode()

	var beatmapset common.BeatmapsetExtended
	if err = bs.Transport.Do(req, &beatmapset); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmapset, nil
}

func (bs *BeatmapsetService) Get(ctx context.Context, beatmapsetId int) (*common.BeatmapsetExtended, error) {
	endpointURL := beatmapsetsBaseAPI + fmt.Sprintf("/%d", beatmapsetId)

	req, err := bs.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	var beatmapset common.BeatmapsetExtended
	if err = bs.Transport.Do(req, &beatmapset); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &beatmapset, nil
}
