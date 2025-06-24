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

func (t *TransactionServiceImpl) DepositMoney(userId int, amount decimal.Decimal) error {
	user, err := t.userService.GetUserByID(userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}
	if err := t.AddCredit(amount, enums.Deposit, user); err != nil {
		log.Printf("Error adding credit: %v", err)
	}
	return nil
}

// FindBy implements TransactionService.
func (t *TransactionServiceImpl) FindBy(column string, value any, single bool) (any, error) {
	if single {
		t.transactionService.FindByOne(column, value)
	}
	return t.transactionService.FindByMany(column, value)
}

// TransferMoneyFromTo implements TransactionService.
func (t *TransactionServiceImpl) TransferMoneyFromTo(fromUserId, toUserId int, amount decimal.Decimal) (model.Transaction, error) {
	if amount.LessThanOrEqual(decimal.Zero) {
		return model.Transaction{}, errors.New("amount must be greater than zero")
	}

	fromUser, err := t.userService.GetUserByID(fromUserId)
	if err != nil {
		return model.Transaction{}, err
	}
	if fromUser == nil {
		return model.Transaction{}, errors.New("from user not found")
	}

	toUser, err := t.userService.GetUserByID(toUserId)
	if err != nil {
		return model.Transaction{}, err
	}
	if toUser == nil {
		return model.Transaction{}, errors.New("to user not found")
	}

	if fromUser.Money.LessThan(amount) {
		return model.Transaction{}, errors.New("insufficient funds")
	}

	fromUser.Money = fromUser.Money.Sub(amount)
	toUser.Money = toUser.Money.Add(amount)

	if err := t.userService.AddCredit(fromUser); err != nil {
		return model.Transaction{}, err
	}

	if err := t.userService.AddCredit(toUser); err != nil {
		return model.Transaction{}, err
	}

	transactionFrom := &model.Transaction{
		UserID: fromUser.ID,
		Amount: amount,
		Status: 0,
		Type:   enums.Transfer,
	}

	if _, err := t.transactionService.CreateTransaction(transactionFrom); err != nil {
		return model.Transaction{}, err
	}

	transactionTo := &model.Transaction{
		UserID: toUser.ID,
		Amount: amount,
		Status: 0,
		Type:   enums.Transfer,
	}
	if _, err := t.transactionService.CreateTransaction(transactionTo); err != nil {
		return model.Transaction{}, err
	}

	return *transactionFrom, nil
}

func (t *TransactionServiceImpl) WithdrawMoney(userId int, amount decimal.Decimal) (model.Transaction, error) {
	if amount.LessThanOrEqual(decimal.Zero) {
		return model.Transaction{}, errors.New("amount must be greater than zero")
	}

	user, err := t.userService.GetUserByID(userId)
	if err != nil {
		return model.Transaction{}, err
	}
	if user == nil {
		return model.Transaction{}, errors.New("user not found")
	}

	if user.Money.LessThan(amount) {
		return model.Transaction{}, errors.New("insufficient funds")
	}

	user.Money = user.Money.Sub(amount)
	if err := t.userService.AddCredit(user); err != nil {
		return model.Transaction{}, err
	}

	tx := &model.Transaction{UserID: userId, Amount: amount.Neg(), Status: enums.Accepted, Type: enums.Withdraw}
	tx, err = t.transactionService.CreateTransaction(tx)
	if err != nil {
		return model.Transaction{}, err
	}

	return *tx, nil
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
		Status: enums.Accepted,
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
