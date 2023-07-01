package models

type Tweets struct {
	CreatedAt uint64 `json:"createdAt"`
	Id        int    `json:"id"`
	Lang      string `json:"lang"`
	Retweet   bool   `json:"retweet"`
	Text      string `json:"text"`
}
