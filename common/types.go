package common

import (
	"fmt"
	"encoding/json"
)

type Ruleset string
const (
	RulesetCatch Ruleset = "fruits"
	RulesetMania Ruleset = "mania"
	RulesetStandard Ruleset = "osu"
	RulesetTaiko Ruleset = "taiko"
)

type RankStatus int
const (
	RankGraveyard RankStatus = iota - 2
	RankWIP 
	RankPending 
	RankRanked 
	RankApproved 
	RankQualified 
	RankLoved 
)

var (
	rankStatusToString = map[RankStatus]string{
		RankGraveyard: "graveyard",
		RankWIP: "wip",
		RankPending: "pending", 
		RankRanked: "ranked", 
		RankApproved: "approved", 
		RankQualified: "qualified", 
		RankLoved: "loved", 
	}

	stringToRankStatus = map[string]RankStatus{
		"graveyard": RankGraveyard,
		"wip": RankWIP,
		"pending": RankPending, 
		"ranked": RankRanked, 
		"approved": RankApproved, 
		"qualified": RankQualified, 
		"loved": RankLoved, 
	}
)

func (rs *RankStatus) String() string {
	val, ok := rankStatusToString[*rs]

	if !ok {
		return ""
	}

	return val
}

func (rs *RankStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if val, ok := stringToRankStatus[str]; ok {
			*rs = val
			return nil
		}
		return fmt.Errorf("invalid status string: %s", str)
	}

	var integer int
	if err := json.Unmarshal(data, &integer); err == nil {
		rankStatus := RankStatus(integer)
		if _, ok := rankStatusToString[rankStatus]; ok {
			*rs = rankStatus
			return nil
		}
		return fmt.Errorf("invalid status int: %d", integer)
	}

	return fmt.Errorf("invalid status format: %s", string(data))
}

type CursorString string
