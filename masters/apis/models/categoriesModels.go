package models

type CategoriesModels struct {
	CategoriesID     string `json:"categories_id"`
	CategoriesDesc   string `json:"categories_desc"`
	CategoriesPrice  string `json:"categories_price"`
	CategoriesStatus string `json:"categories_status"`
	PriceDate        string `json:"price_date"`
}
