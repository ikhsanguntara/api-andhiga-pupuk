package model

type Product struct {
	Model
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Qty         float64 `json:"qty" binding:"qty" binding:"required"`
	Price       float64 `json:"Price" binding:"Price" binding:"required"`
	UOM         string  `json:"uom" binding:"required"`
}
