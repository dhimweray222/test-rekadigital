package web

type CreateTransactionRequest struct {
	CustomerId string `json:"customer_id"`
	Menu       string `json:"menu"`
	Price      int    `json:"price"`
	Qty        int    `json:"qty"`
	Payment    string `json:"payment"`
	Total      int    `json:"total"`
}
