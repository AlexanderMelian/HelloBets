package model

import "github.com/shopspring/decimal"

type User struct {
	ID        int             `json:"id"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Email     string          `json:"email"`
	Money     decimal.Decimal `json:"money"`
	Role      int             `json:"role"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Enabled   bool            `json:"enabled"`
}
