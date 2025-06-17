package model

import "hello_bets/pkg/model/enums"

type Transaction struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	Amount    float64      `json:"amount"`
	Currency  string       `json:"currency"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Status    enums.Status `json:"status"`
	Type      enums.Type   `json:"type"`
}
