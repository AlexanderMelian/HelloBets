package rest

import (
	"hello_bets/pkg/controller"
	"hello_bets/pkg/middleware"
	"hello_bets/pkg/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userController controller.UserController
	userService    service.UserService
}

func NewHandler(s service.UserService, c controller.UserController) (*UserHandler, error) {
	return &UserHandler{userService: s, userController: c}, nil
}

func (h *UserHandler) Routers() *gin.Engine {
	r := gin.Default()
	loginGroup := r.Group("/api/v1/login")
	h.loginRouter(loginGroup)
	userGroup := r.Group("/api/v1/user")
	h.userRouter(userGroup)
	return r

}

func (h *UserHandler) loginRouter(r *gin.RouterGroup) {
	r.POST("/login", h.userController.Login)
}

func (h *UserHandler) userRouter(r *gin.RouterGroup) {
	r.POST("/", h.userController.CreateUser)

	r.Use(middleware.ProtectedHandler())
	r.GET("/:id", h.userController.GetUserByID)
	r.PUT("/:id", h.userController.UpdateUser)
	r.DELETE("/:id", h.userController.DeleteUser)
	r.GET("/", h.userController.FindBy)
}

func StartServer(userController controller.UserController, userSevice service.UserService) {
	handler, err := NewHandler(userSevice, userController)
	if err != nil {
		panic("failed to create handler: " + err.Error())
	}
	r := handler.Routers()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
