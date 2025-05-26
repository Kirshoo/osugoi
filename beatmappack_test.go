package osugoi

import "testing"

func TestGetBeatmapPacks(t *testing.T) {
	_, err := testClient.GetBeatmapPacks()
	if err != nil {
		t.Errorf("GetBeatmapPacks() failed: %v", err)
	}
}

func TestGetBeatmapPack(t *testing.T) {
	packToGet := "S1631"

	pack, err := testClient.GetBeatmapPack(packToGet)
	if err != nil {
		t.Errorf("GetBeatmapPack(%s) failed: %v", packToGet, err)
	}

	t.Logf("%+v\n", pack)
}
