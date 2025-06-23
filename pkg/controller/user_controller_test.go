package controller

import (
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct{}

func GetUsers() []model.User {
	return []model.User{
		{ID: 1, Username: "TestUser", Password: "TestPassword", Email: "test@example.com", Money: decimal.NewFromFloat(100.0)},
		{ID: 2, Username: "TestUser2", Password: "TestPassword2", Email: "test2@example.com", Money: decimal.NewFromFloat(200.0)},
		{ID: 3, Username: "TestUser3", Password: "TestPassword3", Email: "test3@example.com", Money: decimal.NewFromFloat(300.0)},
	}
}

func (m MockUserService) CreateUser(user *dto.UserRequest) (*model.User, error) {
	users := GetUsers()
	for _, u := range users {
		if u.Username == user.Username {
			return nil, nil
		}
	}
	newUser := model.User{
		ID:       len(users) + 1,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Money:    decimal.NewFromFloat(0),
	}
	return &newUser, nil
}

func (m MockUserService) FindBy(column string, value any, single bool) (any, error) {
	users := GetUsers()
	var result []*model.User
	for _, user := range users {
		if column == "username" && user.Username == value {
			result = append(result, &user)
		} else if column == "email" && user.Email == value {
			result = append(result, &user)
		}
	}
	return result, nil
}

func (m MockUserService) GetUserByID(id int) (*model.User, error) {
	users := GetUsers()
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (m MockUserService) UpdateUser(user *dto.UserRequest) (*model.User, error) {
	panic("unimplemented")
}

func (m MockUserService) DeleteUser(id int) error {
	panic("unimplemented")
}

func (m MockUserService) CheckPassword(password, hashPassword string) bool {
	panic("unimplemented")
}

func (m MockUserService) GenerateToken(user *model.User) (string, error) {
	panic("unimplemented")
}

func (m MockUserService) AddCredit(user *model.User) error {
	panic("unimplemented")
}

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := MockUserService{}
	controller, err := NewUserController(mockService)
	if err != nil {
		t.Fatalf("Failed to create user controller: %v", err)
	}

	router := gin.Default()
	router.GET("/users/:id", controller.GetUserByID)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetUserByIDNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := MockUserService{}
	controller, err := NewUserController(mockService)
	if err != nil {
		t.Fatalf("Failed to create user controller: %v", err)
	}

	router := gin.Default()
	router.GET("/users/:id", controller.GetUserByID)
	req, _ := http.NewRequest("GET", "/users/4", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
