package controller

import (
	"fmt"
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/dto"
	"hello_bets/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserControllerInterface
type UserController interface {
	GetUserByID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	FindBy(ctx *gin.Context)
	Login(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(service service.UserService) (*UserControllerImpl, error) {
	if service == nil {
		return nil, fmt.Errorf("user service is nil")
	}
	return &UserControllerImpl{userService: service}, nil
}

func (c *UserControllerImpl) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if id <= 0 {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, user)
}

func (c *UserControllerImpl) CreateUser(ctx *gin.Context) {
	userRequest := &dto.UserRequest{}
	if err := ctx.ShouldBindJSON(userRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	err := validate(*userRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := c.userService.CreateUser(userRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, user)
}
func (c *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	userRequest := &dto.UserRequest{}
	if err := ctx.ShouldBindJSON(userRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if userRequest.ID <= 0 {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	user, err := c.userService.GetUserByID(userRequest.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	userUpdated, err := c.userService.UpdateUser(userRequest)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if userUpdated == nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, user)
}
func (c *UserControllerImpl) FindBy(ctx *gin.Context) {
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

	users, err := c.userService.FindBy(column, value, single)
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

func (c *UserControllerImpl) Login(ctx *gin.Context) {

	LoginRequest := &dto.LoginRequest{}
	if err := ctx.ShouldBindJSON(LoginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}
	if LoginRequest.Username == "" || LoginRequest.Password == "" {
		ctx.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}
	userAny, err := c.userService.FindBy("username", LoginRequest.Username, true)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if userAny == nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	user, ok := userAny.(*model.User)
	if !ok {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if !c.userService.CheckPassword(LoginRequest.Password, user.Password) {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := c.userService.GenerateToken(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

}

func (c *UserControllerImpl) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if id <= 0 {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(204, nil)
}

func validate(u dto.UserRequest) error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
