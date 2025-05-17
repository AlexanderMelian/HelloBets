package model

import "github.com/shopspring/decimal"

type User struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
	Money    decimal.Decimal `json:"money"`
}
