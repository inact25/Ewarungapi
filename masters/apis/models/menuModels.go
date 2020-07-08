package models

type MenuModels struct {
	MenuID     string `json:"menu_id"`
	MenuDesc   string `json:"menu_desc"`
	MenuPrice  string `json:"menu_price"`
	MenuStock  string `json:"menu_stock"`
	MenuStatus string `json:"menu_status"`
	PriceDate  string `json:"price_date"`
}
type MenuPriceModels struct {
	MenuID    string `json:"menu_id"`
	MenuDesc  string `json:"menu_desc"`
	MenuPrice string `json:"menu_price"`
	PriceDate string `json:"price_date"`
}
