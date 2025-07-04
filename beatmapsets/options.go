package beatmapsets

import "strconv"

// BeatmapId option for Lookup has been found
// from aiosu (https://github.com/NiceAesth/aiosu)
// More specifically (https://github.com/NiceAesth/aiosu/blob/master/aiosu/v2/client.py#L1366)
type BeatmapsetOptions struct {
	BeatmapId string `query:"beatmap_id"`
}

type BeatmapsetOption func(*BeatmapsetOptions)
func WithBeatmapId(beatmapId int) BeatmapsetOption {
	return func(options *BeatmapsetOptions) {
		options.BeatmapId = strconv.Itoa(beatmapId)
	}
}
