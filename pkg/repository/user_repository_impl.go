package repository

import (
	"errors"
	"hello_bets/pkg/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &UserRepositoryImpl{db: db}, nil
}

// CreateUser implements UserRepository.
func (u *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser implements UserRepository.
func (u *UserRepositoryImpl) DeleteUser(id int) error {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) FindByMany(column string, value any) ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Where("? = ?", gorm.Expr(column), value).Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindByOne(column string, value any) (*model.User, error) {
	var user model.User
	if err := r.db.Where("? = ?", gorm.Expr(column), value).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) UpdateUser(user *model.User) (*model.User, error) {
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
