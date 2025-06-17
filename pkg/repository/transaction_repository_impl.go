package repository

import (
	"errors"
	"hello_bets/pkg/model"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// CreateTransaction implements TransactionRepository.
func (*TransactionRepositoryImpl) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	panic("unimplemented")
}

// DeleteTransaction implements TransactionRepository.
func (*TransactionRepositoryImpl) DeleteTransaction(id string) error {
	panic("unimplemented")
}

// FindByMany implements TransactionRepository.
func (*TransactionRepositoryImpl) FindByMany(column string, value any) ([]*model.Transaction, error) {
	panic("unimplemented")
}

// FindByOne implements TransactionRepository.
func (*TransactionRepositoryImpl) FindByOne(column string, value any) (*model.Transaction, error) {
	panic("unimplemented")
}

// GetAllTransactions implements TransactionRepository.
func (*TransactionRepositoryImpl) GetAllTransactions() ([]*model.Transaction, error) {
	panic("unimplemented")
}

// GetTransactionByID implements TransactionRepository.
func (*TransactionRepositoryImpl) GetTransactionByID(id string) (*model.Transaction, error) {
	panic("unimplemented")
}

// UpdateTransaction implements TransactionRepository.
func (*TransactionRepositoryImpl) UpdateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	panic("unimplemented")
}

func NewTransactionRepositoryImpl(db *gorm.DB) (TransactionRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &TransactionRepositoryImpl{db: db}, nil
}
