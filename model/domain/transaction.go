package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	Menu       string    `json:"menu"`
	Price      int       `json:"price"`
	Qty        int       `json:"qty"`
	Payment    string    `json:"payment"`
	Total      int       `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
}

func (transaction *Transaction) GenerateID() {
	uuid := uuid.New().String()
	transaction.Id = uuid
}
