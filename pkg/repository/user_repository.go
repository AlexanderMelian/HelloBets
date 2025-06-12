package repository

import (
	"hello_bets/pkg/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	FindByOne(column string, value any) (*model.User, error)
	FindByMany(column string, value any) ([]*model.User, error)
	DeleteUser(id int) error
}
