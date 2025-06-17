package controller

import (
	"fmt"
	"hello_bets/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	FindBy(ctx *gin.Context)
	TransferMoneyFromTo(ctx *gin.Context)
	DepositMoney(ctx *gin.Context)
	WithdrawMoney(ctx *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(service service.TransactionService) (*TransactionControllerImpl, error) {
	if service == nil {
		return nil, fmt.Errorf("user service is nil")
	}
	return &TransactionControllerImpl{transactionService: service}, nil
}

func (c *TransactionControllerImpl) FindBy(ctx *gin.Context) {
	column := ctx.Query("column")
	value := ctx.Query("value")
	uniq := ctx.Query("unique")

	if column == "" || value == "" {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	single, err := strconv.ParseBool(uniq)
	if err != nil {
		single = false
	}

	users, err := c.transactionService.FindBy(column, value, single)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if users == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, users)
}

func (c *TransactionControllerImpl) TransferMoneyFromTo(ctx *gin.Context) {
	// Implement the logic to transfer money from one user to another
	// This is a placeholder for the actual implementation
	ctx.JSON(200, gin.H{"message": "TransferMoneyFromTo method not implemented"})
}
func (c *TransactionControllerImpl) DepositMoney(ctx *gin.Context) {
	// Implement the logic to deposit money into a user's account
	// This is a placeholder for the actual implementation
	ctx.JSON(200, gin.H{"message": "DepositMoney method not implemented"})
}
func (c *TransactionControllerImpl) WithdrawMoney(ctx *gin.Context) {
	// Implement the logic to withdraw money from a user's account
	// This is a placeholder for the actual implementation
	ctx.JSON(200, gin.H{"message": "WithdrawMoney method not implemented"})
}
