package service

import (
	"errors"
	"hello_bets/pkg/configuration"
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/dto"
	"hello_bets/pkg/repository"
	"log"
	"slices"

	"hello_bets/pkg/security"

	"github.com/shopspring/decimal"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	configuration  *configuration.Config
}

func NewUserServiceImpl(configuration *configuration.Config, userRepository repository.UserRepository) (*UserServiceImpl, error) {
	if userRepository == nil {
		return nil, errors.New("user repository is nil")
	}
	if configuration == nil {
		return nil, errors.New("configuration is nil")
	}
	return &UserServiceImpl{userRepository: userRepository, configuration: configuration}, nil

}

func (s *UserServiceImpl) GetUserByID(id int) (*model.User, error) {
	user, err := s.userRepository.FindByOne("id", id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) CreateUser(user *dto.UserRequest) (*model.User, error) {
	if !IsValidEmail(user.Email, s.configuration.PatternMail) {
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
		if !IsValidEmail(user.Email, s.configuration.PatternMail) {
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

func (s *UserServiceImpl) FindBy(column string, value any, single bool) (any, error) {
	err := validateColumn(column)
	if err != nil {
		return nil, err
	}
	if single {
		user, err := s.userRepository.FindByOne(column, value)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := s.userRepository.FindByMany(column, value)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) DeleteUser(id int) error {
	userToDelete, err := s.GetUserByID(id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return errors.New("error try delete user")
	}
	if userToDelete == nil {
		log.Printf("User not found: %d", id)
		return errors.New("error try delete user")
	}
	err = s.userRepository.DeleteUser(id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return errors.New("error try delete user")
	}
	return nil
}

func (s *UserServiceImpl) CheckPassword(password string, HashPassword string) bool {
	loginCounter()
	return CheckPasswordHash(password, HashPassword)
}

func (s *UserServiceImpl) GenerateToken(user *model.User) (string, error) {
	token, err := security.GenerateToken(*user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func validateColumn(column string) error {
	validColumns := []string{"username", "email"}
	if !slices.Contains(validColumns, column) {
		return errors.New("invalid column name")
	}
	return nil
}

func loginCounter() int {

	return 0
}
