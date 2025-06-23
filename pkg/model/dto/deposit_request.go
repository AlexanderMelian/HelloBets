package dto

type DepositRequest struct {
	UserID int     `json:"user_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
