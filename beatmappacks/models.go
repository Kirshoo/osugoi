package beatmappacks

import (
	"fmt"
	"encoding/json"
	"strconv"
	"time"
)

type BeatmapPackTagType struct {
	Name string
	Tag rune
	ModeName string
	Mode rune
}

var (
	tagTypeToNameType = map[rune]string{
		'S': "standard",
		'F': "featured",
		'P': "tournament",
		'L': "loved",
		'R': "chart",
		'T': "theme",
		'A': "artist",
	}

	tagModeToModeName = map[rune]string{
		'C': "osu!catch",
		'M': "osu!mania",
		'T': "osu!taiko",
	}
)

type BeatmapPackTag struct {
	Type BeatmapPackTagType
	Number int
}

func (t *BeatmapPackTag) String() string {
	return fmt.Sprintf("%s%d", string([]rune{t.Type.Tag, t.Type.Mode}), t.Number)
}

func (t *BeatmapPackTag) UnmarshalJSON(data []byte) error {
	var tagString string
	if err := json.Unmarshal(data, &tagString); err != nil {
		return err
	}

	if len(tagString) < 2 {
		return fmt.Errorf("invalid beatmapPack tag: %s", tagString)
	}

	tagRune := rune(tagString[0])
	name, ok := tagTypeToNameType[tagRune]
	if !ok {
		return fmt.Errorf("unknown tag type: %c", tagRune)
	}

	modeRune := rune(tagString[1])
	mode, ok := tagModeToModeName[modeRune]
	if !ok {
		modeRune = '\x00'
		mode = "osu!"
	}

	numIndex := 1
	if ok {
		numIndex += 1
	}

	num, err := strconv.Atoi(tagString[numIndex:])
	if err != nil {
		return fmt.Errorf("convert \"%s\" to number failed: %w", tagString[1:], err)
	}

	packTag := BeatmapPackTag{
		Type: BeatmapPackTagType{
			Tag: tagRune,
			Name: name,
			Mode: modeRune,
			ModeName: mode,
		},
		Number: num,
	}

	*t = packTag
	return nil
}

var (
	StandardPackType BeatmapPackTagType = BeatmapPackTagType{Name: "standard", Tag: 'S'}
	FeaturedArtistPackType BeatmapPackTagType = BeatmapPackTagType{Name: "featured", Tag: 'F'}
	TournamentPackType BeatmapPackTagType = BeatmapPackTagType{Name: "tournament", Tag: 'P'}
	LovedPackType BeatmapPackTagType = BeatmapPackTagType{Name: "loved", Tag: 'L'}
	SpotlightPackType BeatmapPackTagType = BeatmapPackTagType{Name: "chart", Tag: 'R'}
	ThemePackType BeatmapPackTagType = BeatmapPackTagType{Name: "theme", Tag: 'T'}
	ArtistPackType BeatmapPackTagType = BeatmapPackTagType{Name: "artist", Tag: 'A'}
)

type BeatmapPack struct {
	Author string `json:"author"`
	Date time.Time `json:"date"`
	Name string `json:"name"`
	DiffReduction bool `json:"no_diff_reduction"`
	RulesetId int `json:"ruleset_id"` // Possibly change to Ruleset type
	Tag BeatmapPackTag `json:"tag"`
	DownloadUrl string `json:"url"`

	// Optional attributes

	// TODO: Add beatmapset type
	// BeatmapSets *[]Beatmapset `json:"beatmapsets"`
	CompletedBeatmapsetIds []int `json:"user_completion_data.beatmapset_ids"`
	IsCompleted bool `json:"user_completion_data.completed"`
}
