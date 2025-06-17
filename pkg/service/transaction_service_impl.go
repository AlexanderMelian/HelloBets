package service

import (
	"errors"
	"hello_bets/pkg/model"
	"hello_bets/pkg/repository"
)

type TransactionServiceImpl struct {
	transactionService repository.TransactionRepository
}

// DepositMoney implements TransactionService.
func (*TransactionServiceImpl) DepositMoney(userId string, amount float64) (model.Transaction, error) {
	panic("unimplemented")
}

// FindBy implements TransactionService.
func (*TransactionServiceImpl) FindBy(column string, value any, single bool) (any, error) {
	panic("unimplemented")
}

// TransferMoneyFromTo implements TransactionService.
func (*TransactionServiceImpl) TransferMoneyFromTo(fromUserId string, toUserId string, amount float64) (model.Transaction, error) {
	panic("unimplemented")
}

// WithdrawMoney implements TransactionService.
func (*TransactionServiceImpl) WithdrawMoney(userId string, amount float64) (model.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionServiceImpl(transactionService repository.TransactionRepository) (TransactionService, error) {
	if transactionService == nil {
		return nil, errors.New("transaction repository is nil")
	}
	return &TransactionServiceImpl{transactionService: transactionService}, nil
}
