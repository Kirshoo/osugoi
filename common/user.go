package common

import "time"

type KudosuAmount struct {
	Total int64 `json:"total"`
	Available int64 `json:"available"`
}

type User struct {
	AvatarURL string `json:"avatar_url"`
	CountryCode string `json:"country_code"`
	DefaultGroup string `json:"default_group"`
	Id int `json:"id"`
	IsActive bool `json:"is_active"`
	IsBot bool `json:"is_bot"`
	IsDeleted bool `json:"is_deleted"`
	IsOnline bool `json:"is_online"`
	IsSupporter bool `json:"is_supporter"`
	LastVisit time.Time `json:"last_visit"`
	PMFriendsOnly bool `json:"pm_friends_only"`
	ProfileColor string `json:"profile_color"`
	Username string `json:"username"`

	// Optional Attributes
	Cover *interface{} `json:"cover"` // TODO: change to strong type struct
	Country *interface{} `json:"country"` // TODO: change to strong type struct
	IsRestricted *bool `json:"is_restricted,omitempty"`
	Kudosu KudosuAmount `json:"kudosu"`
}

func (u *User) IsPrivatePresence() bool {
	return u.LastVisit.IsZero()
}

type UserExtended struct {
	User

	CoverURL string `json:"cover_url"`
	Discord *string `json:"discord,omitempty"`
	HasSupported bool `json:"has_suported"`
	Interests *string `json:"interests,omitempty"`
	JoinDate time.Time `json:"join_date"`
	Location *string `json:"location"`
	MaxBlocks int64 `json:"max_blocks"`
	MaxFriends int64 `json:"max_friends"`
	Occupation *string `json:"occupation"`
	Playmode Ruleset `json:"playmode"`
	Playstyle []string `json:"playstyle"`
	PostCount int64 `json:"post_count"`
	ProfileHue *int `json:"profile_hue"` // optional between 0 and 359 degrees
	ProfileOrder []string `json:"profile_order"` // TODO: change to enum?
	Title *string `json:"title"`
	TitleURL *string `json:"title_url"`
	Twitter *string `json:"twitter"`
	Website *string `json:"website"`
}
