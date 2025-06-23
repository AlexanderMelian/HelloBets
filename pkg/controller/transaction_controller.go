package controller

import (
	"fmt"
	"hello_bets/pkg/model/dto"
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

func NewTransactionController(transactionService service.TransactionService) (*TransactionControllerImpl, error) {
	if transactionService == nil {
		return nil, fmt.Errorf("user service is nil")
	}
	return &TransactionControllerImpl{transactionService: transactionService}, nil
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

	depositRequest := &dto.DepositRequest{}
	if err := ctx.ShouldBindJSON(depositRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if depositRequest.UserID <= 0 || depositRequest.Amount <= 0 {
		ctx.JSON(400, gin.H{"error": "invalid user ID or amount"})
		return
	}

	///if err := c.transactionService.DepositMoney(depositRequest.UserID, depositRequest.Amount); err != nil {
	//	ctx.JSON(400, gin.H{"error": err.Error()})
	//	return
	//}

	ctx.JSON(200, gin.H{"message": "DepositMoney method not implemented"})
}
func (c *TransactionControllerImpl) WithdrawMoney(ctx *gin.Context) {
	// Implement the logic to withdraw money from a user's account
	// This is a placeholder for the actual implementation
	ctx.JSON(200, gin.H{"message": "WithdrawMoney method not implemented"})
}
