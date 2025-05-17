package controller

import (
	"fmt"
	"hello_bets/intern/model/dto"
	"hello_bets/intern/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
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

func (c *UserController) CreateUser(ctx *gin.Context) {
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
func (c *UserController) UpdateUser(ctx *gin.Context) {
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
