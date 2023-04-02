package model

type Transaction struct {
	Model
	NoTransaction string `json:"no_transaction" `
	Username      string `json:"username" binding:"required"`
	Total         string `json:"Total"`
}
