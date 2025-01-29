package model

type Tweets struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	AccountID string `json:"account_id"`
}
