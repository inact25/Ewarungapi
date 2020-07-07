package models

type CategoryModels struct {
	CategoryID    string `json:"category_id"`
	CategoryDesc  string `json:"category_desc"`
	CategoryDate  string `json:"category_date"`
	CategoryPrice string `json:"category_price"`
}
type CategoryPriceModels struct {
	PriceID   string `json:"price_id"`
	PriceDate string `json:"price_date"`
	Price     string `json:"price"`
}
