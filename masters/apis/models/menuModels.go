package models

type MenuModels struct {
	MenuID    string `json:"menu_id"`
	MenuDesc  string `json:"menu_desc"`
	MenuPrice string `json:"menu_price"`
	MenuStock string `json:"menu_stock"`
}
type MenuPriceModels struct {
	PriceID   string `json:"price_id"`
	PriceDate string `json:"price_date"`
	Price     string `json:"price"`
}
