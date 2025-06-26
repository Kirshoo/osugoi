package events

import (
	"time"
	"github.com/Kirshoo/osugoi/common"
)

type AchievementEvent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Instructions *string `json:"instructions"`
	Mode *common.Ruleset `json:"mode"`
	Ordering int `json:"ordering"`
	Grouping string `json:"grouping"`
	Slug string `json:"slug"`
	IconURL string `json:"icon_url"`
}

type UserEvent struct {
	Username string `json:"username"`
	URL string `json:"url"`
	PreviousUsername *string `json:"previousUsername"`
}

type BeatmapEvent struct {
	Title string `json:"title"`
	URL string `json:"url"`
}

type BeatmapsetEvent struct {
	Title string `json:"title"`
	URL string `json:"url"`
}

type RankEvent struct {
	ScoreRank string `json:"scoreRank"`
	Rank int `json:"rank"`
	Mode common.Ruleset `json:"mode"`
}

type Event struct {
	CreatedAt time.Time `json:"created_at"`
	Id int `json:"id"`

	Achievement *AchievementEvent `json:"achievement"`
	*RankEvent

	Beatmap *BeatmapEvent `json:"beatmap"`
	Beatmapset *BeatmapsetEvent `json:"beatmapset"`
	User *UserEvent `json:"user"`
}
