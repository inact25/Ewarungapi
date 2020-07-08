package models

type ServicesModels struct {
	ServicesID     string `json:"services_id"`
	ServicesDesc   string `json:"services_desc"`
	ServicePrice   string `json:"services_price"`
	ServicesStatus string `json:"services_status"`
	PriceDate      string `json:"price_date"`
}
