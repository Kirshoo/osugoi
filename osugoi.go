package osugoi

import (
	"github.com/Kirshoo/osugoi/auth"
	"github.com/Kirshoo/osugoi/transport"
	"github.com/Kirshoo/osugoi/beatmaps"
	"github.com/Kirshoo/osugoi/beatmappacks"
	"github.com/Kirshoo/osugoi/beatmapsets"
	"github.com/Kirshoo/osugoi/events"
	"github.com/Kirshoo/osugoi/scores"
)

// Represents an API client
// Acts as a main accessing point for various endpoints
// under corresponding namespaces
type Client struct {
	transport *transport.Transport

	Beatmaps *beatmaps.BeatmapService
	BeatmapPacks *beatmappacks.BeatmapPackService
	Beatmapsets *beatmapsets.BeatmapsetService
	Events *events.EventService
	Scores *scores.ScoreService
}

// TODO: Simplify token config creation.
// Ideally user should only import this package.
func NewClient(baseURL string, tokenSource auth.TokenSource, opts ...transport.TransportConfig) *Client {
	trans := transport.New(baseURL, tokenSource, opts...)

	return &Client{
		transport: trans,

		Beatmaps: &beatmaps.BeatmapService{Transport: trans},
		BeatmapPacks: &beatmappacks.BeatmapPackService{Transport: trans},
		Beatmapsets: &beatmapsets.BeatmapsetService{Transport: trans},
		Events: &events.EventService{Transport: trans},
		Scores: &scores.ScoreService{Transport: trans},
	}
}

// Revokes current token. Calls to endpoint methods after revoking the
// token will always result in an error.
func (c *Client) RevokeToken() error {
	return c.transport.RevokeToken()
}
