package service

import "hello_bets/pkg/model"

type TransactionService interface {
	FindBy(column string, value any, single bool) (any, error)
	TransferMoneyFromTo(fromUserId, toUserId string, amount float64) (model.Transaction, error)
	DepositMoney(userId string, amount float64) (model.Transaction, error)
	WithdrawMoney(userId string, amount float64) (model.Transaction, error)
}
