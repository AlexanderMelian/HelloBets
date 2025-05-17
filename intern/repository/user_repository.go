package repository

import (
	"database/sql"
	"hello_bets/intern/model"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	//only return nil because we don't have a database
	return nil, nil
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	//only return nil because we don't have a database
	return nil, nil
}
func (r *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	//only return nil because we don't have a database
	return nil, nil
}
func (r *UserRepository) DeleteUser(id int) error {
	//only return nil because we don't have a database
	return nil
}
func (r *UserRepository) FindBy(column string, value any) ([]*model.User, error) {
	//only return nil because we don't have a database
	return nil, nil
}
