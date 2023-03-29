package model

type Product struct {
	Model
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Stock       float64 `json:"stock" binding:"stock"`
	UOM         string  `json:"UOM" binding:"required"`
	Year        string  `json:"Year"`
}
