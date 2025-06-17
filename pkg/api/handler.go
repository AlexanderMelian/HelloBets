package rest

import (
	"hello_bets/pkg/controller"
	"hello_bets/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	userController        controller.UserController
	transactionController controller.TransactionController
}

func NewHandler(
	userController controller.UserController,
	transactionController controller.TransactionController,
) (*APIHandler, error) {
	return &APIHandler{
		userController:        userController,
		transactionController: transactionController,
	}, nil
}

func (h *APIHandler) Routers() *gin.Engine {
	r := gin.Default()
	loginGroup := r.Group("/api/v1/login")
	h.loginRouter(loginGroup)
	userGroup := r.Group("/api/v1/user")
	h.userRouter(userGroup)
	transferGroup := r.Group("/api/v1/transfer")
	h.transferRouter(transferGroup)
	return r

}

func (h *APIHandler) loginRouter(r *gin.RouterGroup) {
	r.POST("/login", h.userController.Login)
}

func (h *APIHandler) userRouter(r *gin.RouterGroup) {
	r.POST("/", h.userController.CreateUser)

	protected := r.Group("/")
	protected.Use(middleware.ProtectedHandler())
	protected.GET("/:id", h.userController.GetUserByID)
	protected.PUT("/:id", h.userController.UpdateUser)
	protected.DELETE("/:id", h.userController.DeleteUser)
	protected.GET("/", h.userController.FindBy)
}

func (h *APIHandler) transferRouter(r *gin.RouterGroup) {
	protected := r.Group("/")
	protected.Use(middleware.ProtectedHandler())
	protected.POST("/", h.transactionController.DepositMoney)
}

func StartServer(userController controller.UserController, transactionController controller.TransactionController) {
	handler, err := NewHandler(userController, transactionController)
	if err != nil {
		panic("failed to create handler: " + err.Error())
	}
	r := handler.Routers()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
