package repository

import "hello_bets/pkg/model"

type TransactionRepository interface {
	GetTransactionByID(id string) (*model.Transaction, error)
	CreateTransaction(transaction *model.Transaction) (*model.Transaction, error)
	UpdateTransaction(transaction *model.Transaction) (*model.Transaction, error)
	DeleteTransaction(id string) error
	GetAllTransactions() ([]*model.Transaction, error)
	FindByOne(column string, value any) (*model.Transaction, error)
	FindByMany(column string, value any) ([]*model.Transaction, error)
}
