package rest

import (
	"hello_bets/intern/controller"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userController *controller.UserController
}

func (h *UserHandler) Routers() *gin.Engine {
	r := gin.Default()
	//use UserRouter
	userGroup := r.Group("/api/v1/user")
	h.userRouter(userGroup)
	return r

}

func (h *UserHandler) userRouter(r *gin.RouterGroup) {
	r.POST("/", h.userController.CreateUser)

	r.Use()

	r.GET("/:id", h.userController.GetUserByID)
	r.PUT("/:id", h.userController.UpdateUser)
	//r.DELETE("/:id", h.userController.DeleteUser)
	//r.GET("/", h.userController.FindBy)
}
