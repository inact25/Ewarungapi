package models

type ServicesModels struct {
	ServicesID    string `json:"services_id"`
	ServicesDesc  string `json:"services_desc"`
	ServicesPrice string `json:"services_price"`
}
type ServicesPriceModels struct {
	PriceID   string `json:"price_id"`
	PriceDate string `json:"price_date"`
	Price     string `json:"price"`
}
