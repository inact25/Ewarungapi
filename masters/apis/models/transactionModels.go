package models

type TransactionModels struct {
	TransactionID   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
	MenuDesc        string `json:"menu_desc"`
	MenuPrice       string `json:"category_desc"`
	Quantity        string `json:"quantity"`
	CategoryDesc    string `json:"category_desc"`
	FavorPrice      string `json:"favor_price"`
	ServiceDesc     string `json:"service_desc"`
	ServicePrice    string `json:"service_price"`
	SubTotal        string `json:"sub_total"`
}
