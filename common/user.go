package common

import "time"

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
}

func (u *User) IsPrivatePresence() bool {
	return u.LastVisit.IsZero()
}
