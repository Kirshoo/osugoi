package common

import (
	"fmt"
	"encoding/json"

	"github.com/Kirshoo/osugoi/internal/extrafields"
)

type ModSettings struct {
	// AC
	MinimumAccuracy *float64 `json:"minimum_accuracy"`
	Restart *bool `json:"restart"`

	// AD
	// Its a number string ?
	Style *string `json:"style"`

	// BR
	// Number string means enums, right?
	Direction *string `json:"direction"`
	SpinSpeed *float64 `json:"spin_speed"`

	// CL
	AlwaysPlayTailSample *bool `json:"always_play_tail_sample"`
	ClassicHealth *bool `json:"classic_health"`
	ClassicNoteLock *bool `json:"classic_note_lock"`
	FadeHitCircleEarly *bool `json:"fade_hit_circle_early"`
	NoSliderHeadAccuracy *bool `json:"no_slider_head_accuracy"`

	// DA
	DrainRate *float64 `json:"drain_rate"`
	OverallDifficulty *float64 `json:"overall_difficulty"`
	ApproachRate *float64 `json:"approach_rate"`
	CircleSize *float64 `json:"circle_size"`
	ExtendedLimits *bool `json:"extended_limits"`

	// DF
	StartScale *float64 `json:"start_scale"`

	// DP
	MaxDepth *int `json:"max_depth"`

	// DT, NC, HT
	SpeedChange *float64 `json:"speed_change"`

	// DT
	AdjustPitch *bool `json:"adjust_pitch"`

	// EZ
	Retries *int `json:"retries"`

	// HD
	OnlyFadeApproachCircles *bool `json:"only_fade_approach_circles"`

	// MG
	AttractionStrength *float64 `json:"attraction_strength"`

	// MR
	// Also part of enum, i think
	// EDIT: Another string number ?
	Reflection *string `json:"reflection"`

	// MU
	AffectsHitSounds *bool `json:"affects_hit_sounds"`
	EnableMetronome *bool `json:"enable_metronome"`
	InverseMuting *bool `json:"inverse_muting"`
	MuteComboCount *int `json:"mute_combo_count"`

	// NS
	HiddenComboCount *int `json:"hidden_combo_count"`

	// RD
	AngleSharpness *float64 `json:"angle_sharpness"`
	Seed *float64 `json:"seed"`

	// RP
	RepulsionStrength *float64 `json:"repulsion_strength"`

	// SD
	FailOnSliderTail *bool `json:"fail_on_slider_tail"`

	// WG
	Strength *float64 `json:"strength"`

	// WU
	FinalRate *float64 `json:"final_rate"`
	InitialRate *float64 `json:"initial_rate"`

	// Use this if you need to access field
	// that are not yet appointed
	extrafields.WithExtras
}

func (s *ModSettings) UnmarshalJSON(data []byte) error {
	type Alias ModSettings
	var base Alias
	if err := json.Unmarshal(data, &base); err != nil {
		return fmt.Errorf("%w, provided data: %s", err, string(data))
	}

	var all map[string]any
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}

	for field := range extrafields.ExtractKnownFields(&base) {
		delete(all, field)
	}

	*s = ModSettings(base)
	s.ExtraFields = all
	return nil
}

type Mod struct {
	Acronym string `json:"acronym"`
	Settings *ModSettings `json:"settings"`
}
