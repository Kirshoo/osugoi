package common

import "time"

// Collection of links to image used in a cover
type CoverLinks struct {
	Card string `json:"card"`
	Card2x string `json:"card@2x"`
	Cover string `json:"cover"`
	Cover2x string `json:"cover@2x"`
	List string `json:"list"`
	List2x string `json:"list@2x"`
	SlimCover string `json:"slimcover"`
	SlimCover2x string `json:"slimcover@2x"`
}

type Nomination struct {
	BeatmapsetId int `beatmapset_id`
	Reset bool `json:"reset"`
	Rulesets []Ruleset `json:"rulesets"`
	UserId int `json:"user_id"`
}

type Description struct {
	Description string `json:"description"`
}

type Beatmapset struct {
	Artist string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers CoverLinks `json:"covers"`
	Creator string `json:"creator"`
	FavouriteCount int `json:"favourite_count"`
	Id int `json:"id"`
	NSFW bool `json:"nswf"`
	Offset int `json:"offset"`
	PlayCount int `json:"play_count"`
	PreviewURL string `json:"preview_url"`
	Source string `json:"source"`
	Status RankStatus `json:"status"`
	Spotlight bool `json:"spotlight"`
	Title string `json:"title"`
	TitleUnicode string `json:"title_unicode"`
	UserId int `json:"user_id"`
	Video bool `json:"video"`

	// Optional Fields

	Beatmaps []BeatmapExtended `json:"beatmaps"`
	// TODO: Make Converted beatmap struct
	Converts []interface{} `json:"convers"`
	CurrentNominations []Nomination `json:"current_nominations"`
	CurrentUserAttributes *interface{} `json:"current_user_attributes"`
	Description Description `json:"description"`

	// Docs mention discussions field, but haven't seen
	// it while crawling yet. Need more testing...
	// Discussions *interface{} `json:"discussions"`
	
	// Docs mention events field, but haven't seen
	// it while crawling yet. Need more testing...
	// Events *interface{} `json:"events"`

	GenreId int `json:"genre_id"`
	LanguageId int `json:"language_id"`

	// Docs have has_favourite as an optional field, but I
	// was unable to get it during crawling

	// Only available when requesting from /beatmapsets endpoint
	PackTags *[]string `json:"pack_tags"`

	// Ommited when requested from /users/{id}/beatmapsets/
	Ratings []int `json:"ratings"`
	RecentFavourites []User `json:"recent_favourites"`
	RelatedUsers []User `json:"related_users"`
	User *User `json:"user"`

	TrackId *int `json:"track_id"`
}

type Availability struct {
	DownloadDisabled bool `json:"download_disabled"`
	MoreInformation *string `json:"more_information"`
}

type Genre struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Language struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Hype struct {
	Current int `json:"current"`
	Required int `json:"required"`
}

type RequiredNominations struct {
	MainRuleset int `json:"main_ruleset"`
	OtherRuleset int `json:"non_main_ruleset"`
}

type NominationsSummary struct {
	Current int `json:"current"`
	EligibleMainRulesets *[]Ruleset `json:"eligible_main_rulesets"`
	Required RequiredNominations `json:"required_meta"`
}

type BeatmapsetExtended struct {
	Beatmapset

	Availability Availability `json:"availability"`
	BPM float64 `json:"bpm"`
	CanBeHyped bool `json:"can_be_hyped"`
	DeletedAt time.Time `json:"deleted_at"`
	DiscussionEnabled bool `json:"discussion_enabled"`
	DiscussionLocked bool `json:"discussion_locked"`
	Hype *Hype `json:"hype"`
	IsScorable bool `json:"is_scorable"`
	LastUpdated time.Time `json:"last_updated"`
	LegacyThreadURL *string `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary `json:"nominations_summary"`

	// Redundant, because it's is exactly the same 
	// as Status of Base Beatmapset
	RankedStatus RankStatus `json:"ranked"`

	RankedAt time.Time `json:"ranked_date"`
	Rating float64 `json:"rating"`

	// Docs mention source again. I'll ommit because 
	// its already part of Base Beatmapset
	// Source string `json:"source"`

	HasStoryboard bool `json:"storyboard"`
	SubmittedAt time.Time `json:"submitted_date"0
	`
	// Maybe should be changed to []string and
	// split tags by ' ' (space)
	Tags string `json:"tags"`

	// Ommited when requested from /users/{id}/beatmapsets
	Genre *Genre `json:"genre"`
	Language *Language `json:"language"`
}

func (b *BeatmapsetExtended) IsDeleted() bool {
	return !b.DeletedAt.IsZero()
}

func (b *BeatmapsetExtended) IsRanked() bool {
	return !b.RankedAt.IsZero()
}
