package service

import (
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/enums"

	"github.com/shopspring/decimal"
)

type TransactionService interface {
	FindBy(column string, value any, single bool) (any, error)
	TransferMoneyFromTo(fromUserId, toUserId int, amount decimal.Decimal) (model.Transaction, error)
	DepositMoney(userId int, amount decimal.Decimal) error
	WithdrawMoney(userId int, amount decimal.Decimal) (model.Transaction, error)
	AddCredit(amout decimal.Decimal, operationType enums.Type, userId *model.User) error
}
