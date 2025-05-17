package service

import (
	"hello_bets/intern/model"
	"hello_bets/intern/model/dto"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *dto.UserRequest) (*model.User, error)
	UpdateUser(user *dto.UserRequest) (*model.User, error)
	FindBy(column string, value any) ([]*model.User, error)
}
