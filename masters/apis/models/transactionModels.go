package models

type TransactionModels struct {
	TransactionID   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	ServicesDesc    string `json:"services_desc"`
	ServicePrice    string `json:"service_price"`
	MenuDesc        string `json:"menu_desc"`
	MenuPrice       string `json:"menu_price"`
	CategoryDesc    string `json:"category_desc"`
	CategoryPrice   string `json:"category_price"`
	Qty             string `json:"qty"`
	SubTotal        string `json:"sub_total"`
}
