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
func (r *TransactionRepositoryImpl) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	if transaction == nil {
		return nil, errors.New("transaction is nil")
	}
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// DeleteTransaction implements TransactionRepository.
func (r *TransactionRepositoryImpl) DeleteTransaction(id string) error {
	if err := r.db.Delete(&model.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}

// FindByMany implements TransactionRepository.
func (r *TransactionRepositoryImpl) FindByMany(column string, value any) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.db.Where("? = ?", gorm.Expr(column), value).Find(&transactions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []*model.Transaction{}, nil
		}
		return nil, err
	}
	return transactions, nil
}

// FindByOne implements TransactionRepository.
func (r *TransactionRepositoryImpl) FindByOne(column string, value any) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.db.Where("? = ?", gorm.Expr(column), value).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

// GetAllTransactions implements TransactionRepository.
func (r *TransactionRepositoryImpl) GetAllTransactions() ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionByID implements TransactionRepository.
func (r *TransactionRepositoryImpl) GetTransactionByID(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

// UpdateTransaction implements TransactionRepository.
func (r *TransactionRepositoryImpl) UpdateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	if transaction == nil {
		return nil, errors.New("transaction is nil")
	}
	if err := r.db.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionRepositoryImpl(db *gorm.DB) (TransactionRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &TransactionRepositoryImpl{db: db}, nil
}
