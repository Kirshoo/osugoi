package beatmaps

import "github.com/Kirshoo/osugoi/common"

type OsuDifficultyAttributes struct {
	AimDifficulty *float64 `json:"aim_difficulty"`
	AimDifficultSliderCount *float64 `json:"aim_difficult_slider_count"`
	AimDifficultStrainCount *float64 `json:"aim_difficult_strain_count"`
	SpeedDifficulty *float64 `json:"speed_difficulty"`
	SpeedNodeCount *float64 `json:"speed_node_count"`
	SpeedDifficultStrainCount *float64 `json:"speed_difficult_strain_count"`
	SliderFactor *float64 `json:"slider_factor"`
}

type TaikoDifficultyAttributes struct {
	MonoStaminaFactor *float64 `json:"mono_stamina_factor"`
}

type DifficultyAttributes struct {
	StarRating float64 `json:"star_rating"`
	MaxCombo int `json:"max_combo"`

	OsuDifficultyAttributes
	TaikoDifficultyAttributes
}

type BeatmapPlaycount struct {
	BeatmapId int `json:"beatmap_id"`
	Beatmap *common.Beatmap `json:"beatmap"`
	//Beatmapset *common.Beatmapset `json:"beatmapset"`
	Count int `json:"count"`
}

type BeatmapUserScore struct {
	Position int `json:"position"`
	Score common.Score `json:"score"`
}

type BeatmapScores struct {
	Scores []common.Score `json:"scores"`

	// Will be moved to user_score
	UserScore *BeatmapUserScore `json:"userScore"` 
}
