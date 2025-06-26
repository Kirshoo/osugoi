package common

import "time"

type Failtimes struct {
	Exit *[]int `json:"exit"`
	Fail *[]int `json:"fail"`
}

type BeatmapOwner struct {
	Id int `json:"id"`
	Username string `json:"username"`
}

type Beatmap struct {
	SetId int `json:"beatmapset_id"`
	Difficulty float64 `json:"difficulty_rating"`
	Id int `json:"id"`
	Mode Ruleset `json:"mode"`
	Status string `json:"status"`
	Length int `json:"total_length"`
	UserId int `json:"user_id"`
	Version string `json:"version"`

	// Optional Attributes
	// TODO:
	// Issue: is flexible. Can either be Beatmapset, BeatmapsetExtended or null!
	// Beatmapset *Beatmapset `json:"beatmapset"`
	Checksum *string `json:"checksum"`
	CurrentUserPlaycount *int `json:"current_user_playcount"`
	Failtimes *Failtimes `json:"failtimes"`
	MaxCombo *int `json:"max_combo"`
	Owners *[]BeatmapOwner `json:"owners"`
}

type BeatmapExtended struct {
	Beatmap

	BPM *float64 `json:"bpm"`
	AR float64 `json:"ar"`
	CS float64 `json:"cs"`
	Drain float64 `json:"drain"` //HP?
	Accuracy float64 `json:"accuracy"` //OD?
	HitLength int `json:"hit_length"` //OD?

	CircleCount int `json:"count_circles"`
	SliderCount int `json:"count_slider"`
	SpinnerCount int `json:"count_spinner"`

	BeatmapsetId int `json:"beatmapset_id"`
	DeletedAt *time.Time `json:"deleted_at"`
	LastUpdatedAt time.Time `json:"last_updated"`
	IsScorable bool `json:"is_scorable"`
	ModeInt int `json:"mode_int"`
	PassCount int `json:"passcount"`
	PlayCount int `json:"playcount"`
	RankedStatus RankStatus `json:"ranked"`
	URL string `json:"url"`

	// Need more testing to determine purpose
	Convert bool `json:"convert"`
}

