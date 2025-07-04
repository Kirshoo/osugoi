package beatmappacks

import (
	"fmt"
	"net/http"
	"context"

	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/common"
	"github.com/Kirshoo/osugoi/internal/options"
	"github.com/Kirshoo/osugoi/internal/optionquery"
)

const beatmapPacksBaseAPI string = "/api/v2/beatmaps/packs"

type BeatmapPackService struct {
	Transport *transport.Transport
}

// Returns error only when option parameter is
// passed as nil, otherwise never errors
func assignParameters(opts []BeatmapPackOption, option *BeatmapPackOptions) error {
	if option == nil {
		return fmt.Errorf("BeatmapOptions reference cannot be nil")
	}

	for _, opt := range opts {
		opt(option)
	}

	return nil
}

func (bp *BeatmapPackService) Get(ctx context.Context, packTag string, opts ...BeatmapPackOption) (*common.BeatmapPack, error) {
	endpointURL := fmt.Sprintf(beatmapPacksBaseAPI + "/%s", packTag)
	allowedParameters := []string{"legacy_only"}

	req, err := bp.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapPackOptions{}
	assignParameters(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Accept", "application/json")

	var pack common.BeatmapPack
	if err = bp.Transport.Do(req, &pack); err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return &pack, nil
}

func (bp *BeatmapPackService) List(ctx context.Context, opts ...BeatmapPackOption) (*[]common.BeatmapPack, common.CursorString, error) {
	endpointURL := beatmapPacksBaseAPI
	allowedParameters := []string{"type", "cursor_string"}

	req, err := bp.Transport.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, "", fmt.Errorf("creating request: %w", err)
	}

	parameters := BeatmapPackOptions{}
	assignParameters(opts, &parameters)
	query := optionquery.Convert(parameters)

	options.Filter(allowedParameters, &query)
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Accept", "application/json")

	var list struct {
		Packs []common.BeatmapPack `json:"beatmap_packs"`
		Cursor common.CursorString `json:"cursor_string"`
	}
	if err = bp.Transport.Do(req, &list); err != nil {
		return nil, "", fmt.Errorf("performing request: %w", err)
	}

	return &list.Packs, list.Cursor, nil
}
