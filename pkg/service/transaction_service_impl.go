package service

import (
	"errors"
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/enums"
	"hello_bets/pkg/repository"
	"log"

	"github.com/shopspring/decimal"
)

type TransactionServiceImpl struct {
	transactionService repository.TransactionRepository
	userService        UserService
}

// DepositMoney implements TransactionService.
func (t *TransactionServiceImpl) DepositMoney(userId int, amount decimal.Decimal) error {
	user, err := t.userService.GetUserByID(userId)

	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	AddCreditErr := t.AddCredit(amount, enums.Deposit, user)
	if AddCreditErr != nil {
		log.Println("Error adding credit:", AddCreditErr)
		return AddCreditErr
	}

	transaction := model.Transaction{
		UserID: userId,
		Amount: amount,
		Status: 0,
		Type:   0,
	}

	transactionPtr, err := t.transactionService.CreateTransaction(&transaction)
	if err != nil {
		return err
	}
	transaction = *transactionPtr
	log.Println("Transaction created:", transaction)
	return nil
}

// FindBy implements TransactionService.
func (*TransactionServiceImpl) FindBy(column string, value any, single bool) (any, error) {
	panic("unimplemented")
}

// TransferMoneyFromTo implements TransactionService.
func (*TransactionServiceImpl) TransferMoneyFromTo(fromUserId, toUserId int, amount decimal.Decimal) (model.Transaction, error) {
	panic("unimplemented")
}

// WithdrawMoney implements TransactionService.
func (*TransactionServiceImpl) WithdrawMoney(userId int, amount decimal.Decimal) (model.Transaction, error) {
	panic("unimplemented")
}

func (t *TransactionServiceImpl) AddCredit(amount decimal.Decimal, operationType enums.Type, user *model.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	user.Money = user.Money.Add(amount)

	err := t.userService.AddCredit(user)
	if err != nil {
		return err
	}

	transaction := &model.Transaction{
		UserID: user.ID,
		Amount: amount,
		Status: 0,
		Type:   operationType,
	}

	transaction, err = t.transactionService.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	log.Printf("Transaction added: %+v", transaction)
	return nil
}

func NewTransactionServiceImpl(transactionService repository.TransactionRepository, userService UserService) (TransactionService, error) {
	if transactionService == nil {
		return nil, errors.New("transaction repository is nil")
	}
	if userService == nil {
		return nil, errors.New("user service is nil")
	}
	return &TransactionServiceImpl{transactionService: transactionService, userService: userService}, nil
}
