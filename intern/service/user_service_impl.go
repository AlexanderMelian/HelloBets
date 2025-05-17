package service

import (
	"errors"
	"hello_bets/intern/model"
	"hello_bets/intern/model/dto"
	"hello_bets/intern/repository"
	"log"

	"github.com/shopspring/decimal"
)

type UserServiceImpl struct {
	userRepository *repository.UserRepository
}

func NewUserServiceImpl(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) GetUserByID(id int) (*model.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *dto.UserRequest) (*model.User, error) {
	if !IsValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}
	if IsValidPassword(user.Password) {
		return nil, errors.New("invalid password format")
	}
	passwordHashed, err := HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password: " + err.Error())
	}
	newUser := &model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: passwordHashed,
		Money:    decimal.NewFromInt(0),
	}
	createdUser, err := s.userRepository.CreateUser(newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, errors.New("failed to create user")
	}
	return createdUser, nil
}

func (s *UserServiceImpl) UpdateUser(user *dto.UserRequest) (*model.User, error) {
	if user.ID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	userToUpdate, err := s.GetUserByID(user.ID)
	if err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}
	if userToUpdate == nil {
		return nil, errors.New("user not found")
	}
	if user.Email != "" {
		if !IsValidEmail(user.Email) {
			return nil, errors.New("invalid email format")
		}
		userToUpdate.Email = user.Email
	}
	if user.Password != "" {
		if !IsValidPassword(user.Password) {
			return nil, errors.New("invalid password format")
		}
		passwordHashed, err := HashPassword(user.Password)
		if err != nil {
			return nil, errors.New("failed to hash password: " + err.Error())
		}
		userToUpdate.Password = passwordHashed
	}
	s.userRepository.UpdateUser(userToUpdate)
	return userToUpdate, nil
}

func (s *UserServiceImpl) FindBy(column string, value any) ([]*model.User, error) {
	user, err := s.userRepository.FindBy(column, value)
	if err != nil {
		return nil, err
	}
	return user, nil
}
