package common

import "time"

type MultiplayerScores struct {
	CursorString CursorString `json:"cursor_string"`
	Params interface{} `json:"params"`
	Scores []Score `json:"scores"`
	Total *int `json:"total"`
	UserScore *Score `json:"user_score"`
}

type MultiplayerScoresAround struct {
	Higher MultiplayerScores `json:"higher"`
	Lower MultiplayerScores `json:"lower"`
}

// TODO: Figure out what these are responsible for
type MaximumStatistics struct {
	Great int `json:"great"`
	IgnoreHit int `json:"ignore_hit"`
	LargeBonus int `json:"large_bonus"`
	LegacyComboIncrease int `json:"legacy_combo_increase"`
	Perfect int `json:"perfect"`
	SliderTailHit int `json:"slider_tail_hit"`
	SmallBonus int `json:"small_bonus"`
	SmallTickHit int `json:"small_tick_hit"`
}

type Statistics struct {
	ComboBreak int `json:"combo_break"`
	Good int `json:"good"`
	Great int `json:"great"`
	IgnoreHit int `json:"ignore_hit"`
	IgnoreMiss int `json:"ignore_miss"`
	LargeBonus int `json:"large_bonus"`
	LargeTickHit int `json:"large_tick_hit"`
	LargeTickMiss int `json:"large_tick_miss"`
	Meh int `json:"meh"`
	Ok int `json:"ok"`
	Perfect int `json:"perfect"`
	SliderTailHit int `json:"slider_tail_hit"`
	SmallBonus int `json:"small_bonus"`
	SmallTickHit int `json:"small_tick_hit"`
	SmallTickMiss int `json:"small_tick_miss"`
}

type Score struct {
	Accuracy float32 `json:"accuracy"`
	BeatmapId int `json:"beatmap_id"`
	// TODO: Investigate BestID and BuildID
	BestId *int `json:"best_id"`
	BuildId *int `json:"build_id"`
	StartedAt *time.Time `json:"started_at"`
	EndedAt time.Time `json:"ended_at"`
	HasReplay bool `json:"has_replay"`
	Id int `json:"id"`
	IsPerfectCombo bool `json:"is_perfect_combo"`
	MaxCombo int `json:"max_combo"`
	Passed bool `json:"passed"`
	PP *float32 `json:"pp"`
	Rank string `json:"rank"`
	RulesetId int `json:"ruleset_id"` // May need to convert to string later
	TotalScore int `json:"total_score"`
	Type string `json:"type"`
	UserId int `json:"user_id"`
	
	// Only used for multiplayer score
	PlaylistItemId int `json:"playlist_item_id"`
	RoomId int `json:"room_id"`

	// Potentially depricated fields due to being for solo-score type
	ClassicTotalScore int `json:"classic_total_score"`
	Preserve bool `json:"preserve"`
	Processed bool `json:"processed"`
	Ranked bool `json:"ranked"`

	// Currently not documented types :/
	MaximumStatistics MaximumStatistics `json:"maximum_statistics"`
	Mods []Mod `json:"mods"`
	Statistics Statistics `json:"statistics"`
	
	// Legacy fields
	LegacyPerfect bool `json:"legacy_perfect"`
	LegacyScoreId *int `json:"legacy_score_id"`
	LegacyTotalScore int `json:"legacy_total_score"`

	// Optional Attributes
	Beatmap *Beatmap `json:"beatmap"`
	BeatmapSet *interface{} `json:"beatmapset"`
	CurrentUserAttributes *interface{} `json:"current_user_attributes"`
	Match *interface{} `json:"match"` // Only for legacy match score
	Position *int `json:"position"` // Only for multiplayer score
	RankCountry *interface{} `json:"rank_country"`
	RankGlobal *interface{} `json:"rank_global"`
	// Only for multiplayer score
	ScoresAround *MultiplayerScoresAround `json:"scores_around"`
	User *interface{} `json:"user"`
	Weight *interface{} `json:"weight"`
}
