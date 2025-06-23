package integration

import (
	"hello_bets/pkg/configuration"
	"hello_bets/pkg/controller"
	"hello_bets/pkg/model"
	"hello_bets/pkg/model/dto"
	"hello_bets/pkg/repository"
	"hello_bets/pkg/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestUserIntegration(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	var cfg = &configuration.Config{
		PatternMail: "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
	}
	// Inicializar repositorio, servicio y controlador
	repo, err := repository.NewUserRepository(db)
	assert.NoError(t, err)
	svc, err := service.NewUserServiceImpl(cfg, repo) // Puedes pasar una configuración simulada si es necesario
	assert.NoError(t, err)
	ctrl, err := controller.NewUserController(svc)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", ctrl.CreateUser)
	router.GET("/users/:id", ctrl.GetUserByID)

	// Prueba de creación de usuario
	t.Run("CreateUser", func(t *testing.T) {
		userRequest := dto.UserRequest{
			Username: "testuser",
			Password: "Password123",
			Email:    "test@example.com",
		}

		reqBody := `{"username":"testuser","password":"Password123","email":"test@example.com"}`
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

		// Verificar que el usuario se creó en la base de datos
		var user model.User
		err := db.First(&user, "username = ?", userRequest.Username).Error
		assert.NoError(t, err)
		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "testuser")
	})

	t.Run("GetUserByIDNotFound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/users/999", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
