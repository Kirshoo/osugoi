package events

import (
	"fmt"
	"context"
	"net/http"

	"github.com/Kirshoo/osugoi/client"
	"github.com/Kirshoo/osugoi/common"

	"github.com/Kirshoo/osugoi/internal/optionquery"
)

const baseEventAPI string = "/api/v2/events"

type EventService struct {
	Client *client.Client
}

type listEventResponse struct {
	Events []Event `json:"events"`
	Cursor common.CursorString `json:"cursor_string"`
}

func (e *EventService) List(ctx context.Context, opts ...EventOption) (*[]Event, common.CursorString, error) {
	endpointURL := baseEventAPI

	req, err := e.Client.NewRequest(ctx, http.MethodGet, endpointURL, nil)
	if err != nil {
		return nil, "", fmt.Errorf("creating request: %w", err)
	}

	var options EventOptions
	for _, opt := range opts {
		opt(&options)
	}
	query := optionquery.Convert(options)

	req.URL.RawQuery = query.Encode()
	req.Header.Add("Accept", "application/json")

	var response listEventResponse
	if err = e.Client.Do(req, &response); err != nil {
		return nil, "", fmt.Errorf("performing request: %w", err)
	}

	return &response.Events, response.Cursor, nil
}
