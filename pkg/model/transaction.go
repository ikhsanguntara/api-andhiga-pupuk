package model

type Transaction struct {
	Model
	Code     string `json:"code" binding:"required"`
	Username string `json:"username" binding:"required"`
	Total    string `json:"Total"`
}
