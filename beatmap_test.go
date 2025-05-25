package osugoi

import (
	"testing"
)

// https://osu.ppy.sh/beatmapsets/2349118#osu/5054112
const testBeatmapId = 5054112

func TestGetBeatmap(t *testing.T) {
	_, err := testClient.GetBeatmap(testBeatmapId)
	if err != nil {
		t.Errorf("error getting beatmap %d: %v",
			testBeatmapId, err)
	}
}
