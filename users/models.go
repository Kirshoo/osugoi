package users

import "time"

type KudosuAction string
const (
	KudosuGiveAction KudosuAction = "give"
	KudosuVoteGiveAction = "vote.give"
	KudosuResetAction = "reset"
	KudosuVoteResetAction = "vote.reset"
	KudosuRevokeAction = "revoke"
	KudosuVoteRevokeAction = "vote.revoke"
)

type KudosuGiver struct {
	URL string `json:"url"`
	Username string `json:"username"`
}

type KudosuPost struct {
	URL *string `json:"url"`
	Title *string `json:"title"`
}

type KudosuHistory struct {
	Id int `json:"id"`
	Action KudosuAction `json:"action"`
	Amount int `json:"amount"`

	// TODO: transform into an Enum
	Model string `json:"model"`

	CreatedAt time.Time `json:"created_at"`
	Giver *KudosuGiver `json:"giver"`
	Post KudosuPost `json:"post"`
}
