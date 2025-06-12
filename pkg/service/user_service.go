package service

import (
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/dto"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *dto.UserRequest) (*model.User, error)
	UpdateUser(user *dto.UserRequest) (*model.User, error)
	FindBy(column string, value any, single bool) (any, error)
	DeleteUser(id int) error
	CheckPassword(password, HashPassword string) bool
	GenerateToken(user *model.User) (string, error)
}
