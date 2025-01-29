package model

type Tweet struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	AccountID string `json:"account_id"`
}
