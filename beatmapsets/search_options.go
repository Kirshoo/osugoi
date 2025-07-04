package beatmapsets

import (
	"fmt"
	"strings"
	"time"
)

type BeatmapStatus string
const (
	StatusRanked BeatmapStatus = "ranked"
	StatusApproved = "approved"
	StatusPending = "pending"
	StatusNotSubmitted = "notsubmitted"
	StatusUnknown = "unknown"
	StatusLoved = "loved"
)

type SubqueryType int
const (
	SubqueryKeyword SubqueryType = iota
	SubqueryValueString
	SubqueryValueNumber
)

type SubqueryParameter struct {
	Type SubqueryType
	Key string
	Value interface{}
	Operator SearchOperator
}

func BuildSubquery(parameters ...SubqueryParameter) string {
	var parts []string

	for _, parameter := range parameters {
		switch parameter.Type {
		case SubqueryKeyword:
			if val, ok := parameter.Value.(string); ok {
				parts = append(parts, val)
			}
		case SubqueryValueString:
			if val, ok := parameter.Value.(string); ok {
				parts = append(parts, fmt.Sprintf(`%s%s""%s""`, 
					parameter.Key, parameter.Operator.String(), val))
			}
		case SubqueryValueNumber:
			switch val := parameter.Value.(type) {
			case int:
				parts = append(parts, fmt.Sprintf("%s%s%d",
					parameter.Key, parameter.Operator.String(), val))
			case float64:
				parts = append(parts, fmt.Sprintf("%s%s%.2f",
					parameter.Key, parameter.Operator.String(), val))
			}
		}
	}

	return strings.Join(parts, " ")
}

// General query by the specified keyword
func WithKeyword(keyword string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryKeyword,
		Key: "",
		Value: keyword,
		Operator: OpNoop,
	}
}

// Query for sets with song(s) by specified artist
func WithArtist(artistName string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "artist",
		Value: artistName,
		Operator: OpEquals,
	}
}

// Query for sets that have this inside their title
func WithTitle(title string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "title",
		Value: title,
		Operator: OpEquals,
	}
}

// Query for sets that are in close relation to this medium
func WithSource(source string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "source",
		Value: source,
		Operator: OpEquals,
	}
}

// Query for sets that include songs by specified featured artist identifier
func WithFeaturedArtistId(artistId int) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "featured_artist",
		Value: artistId,
		Operator: OpEquals,
	}
}

// Query for sets created by specified username
func WithCreator(username string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "creator",
		Value: username,
		Operator: OpEquals,
	}
}

// Query for sets that have beatmaps with specified difficulty name
func WithDifficultyName(diffName string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "difficulty",
		Value: diffName,
		Operator: OpEquals,
	}
}

// Query by set's beatmap approach rate parameter
func WithApproachRate(rate float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "ar",
		Value: rate,
		Operator: OpEquals,
	}
}

// Alias for WithApproachRate(float64)
func WithAR(ar float64) SubqueryParameter {
	return WithApproachRate(ar)
}

// Query by set's beatmap circle size parameter
func WithCircleSize(size float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "cs",
		Value: size,
		Operator: OpEquals,
	}
}

// Alias for WithCircleSize(float64)
func WithCS(cs float64) SubqueryParameter {
	return WithCircleSize(cs)
}

// Query by set's beatmap accuracy parameter
func WithAccuracy(accuracy float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "od",
		Value: accuracy,
		Operator: OpEquals,
	}
}

// Alias for WithAccuracy(float64)
func WithOD(od float64) SubqueryParameter {
	return WithAccuracy(od)
}

// Query by set's beatmap drain rate parameter
func WithDrainRate(rate float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "hp",
		Value: rate,
		Operator: OpEquals,
	}
}

// Alias for WithDrainRate(float64)
func WithHP(hp float64) SubqueryParameter {
	return WithDrainRate(hp)
}

// Query by set's beatmap star difficulty
func WithStarDifficulty(stars float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "stars",
		Value: stars,
		Operator: OpEquals,
	}
}

// Query by set's beatmap beats per minute
func WithBPM(bpm float64) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "bpm",
		Value: bpm,
		Operator: OpEquals,
	}
}

// Query by set's beatmap length in seconds
func WithLength(seconds int) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "length",
		Value: seconds,
		Operator: OpEquals,
	}
}

// Query by set's beatmap circle count
// This looks for sets, that includes a beatmap
// with a specified counter
func WithCircleCount(count int) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "circles",
		Value: count,
		Operator: OpEquals,
	}
}

// Query by set's beatmap slider count
// This looks for sets, that includes a beatmap
// with a specified counter
func WithSliderCount(count int) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueNumber,
		Key: "sliders",
		Value: count,
		Operator: OpEquals,
	}
}

// Query by set's beatmap status
func WithStatus(status BeatmapStatus) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "status",
		Value: status,
		Operator: OpEquals,
	}
} 

// Query by timestamp when the map was created/uploaded
func WithCreationDate(t time.Time) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "created",
		Value: t.Format(time.DateOnly),
		Operator: OpEquals,
	}
}

// Query by timestamp when the map was last updated
func WithUpdateDate(t time.Time) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "updated",
		Value: t.Format(time.DateOnly),
		Operator: OpEquals,
	}
}

// Query by timestamp when the map was ranked
func WithRankingDate(t time.Time) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "ranked",
		Value: t.Format(time.DateOnly),
		Operator: OpEquals,
	}
}

// Query with a specific user tag
func WithTag(tag string) SubqueryParameter {
	return SubqueryParameter{
		Type: SubqueryValueString,
		Key: "tag",
		Value: tag,
		Operator: OpEquals,
	}
}
