package model

import (
	"hello_bets/pkg/model/enums"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        int             `json:"id"`
	UserID    int             `json:"user_id"`
	Amount    decimal.Decimal `json:"amount"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Status    enums.Status    `json:"status"`
	Type      enums.Type      `json:"type"`
}
