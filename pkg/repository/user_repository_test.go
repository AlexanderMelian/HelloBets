package repository

import (
	"hello_bets/pkg/model"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	// Usar SQLite en memoria para pruebas
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar el modelo de usuario
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo, err := NewUserRepository(db)
	assert.NoError(t, err)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	createdUser, err := repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "testuser", createdUser.Username)
	assert.Equal(t, "test@example.com", createdUser.Email)
}

func TestFindByOne(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo, err := NewUserRepository(db)
	assert.NoError(t, err)

	// Crear un usuario de prueba
	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	_, err = repo.CreateUser(user)
	assert.NoError(t, err)

	foundUser, err := repo.FindByOne("username", "testuser")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "testuser", foundUser.Username)
	assert.Equal(t, "test@example.com", foundUser.Email)
	assert.Equal(t, decimal.NewFromInt(0), foundUser.Money)
}

func TestDeleteUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo, err := NewUserRepository(db)
	assert.NoError(t, err)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	createdUser, err := repo.CreateUser(user)
	assert.NoError(t, err)

	err = repo.DeleteUser(createdUser.ID)
	assert.NoError(t, err)

	foundUser, err := repo.FindByOne("id", createdUser.ID)
	assert.NoError(t, err)
	assert.Nil(t, foundUser)
}

func TestFindByMany(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo, err := NewUserRepository(db)
	assert.NoError(t, err)

	users := []*model.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1"},
		{Username: "user2", Email: "user2@example.com", Password: "password2"},
		{Username: "user3", Email: "user3@example.com", Password: "password3"},
	}
	for _, user := range users {
		_, err := repo.CreateUser(user)
		assert.NoError(t, err)
	}

	foundUsers, err := repo.FindByMany("email", "user1@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, foundUsers)
	assert.Len(t, foundUsers, 1)
	assert.Equal(t, "user1", foundUsers[0].Username)
}

func TestFindByManyReturnNil(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo, err := NewUserRepository(db)
	assert.NoError(t, err)

	users := []*model.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1"},
		{Username: "user2", Email: "user2@example.com", Password: "password2"},
		{Username: "user3", Email: "user3@example.com", Password: "password3"},
	}
	for _, user := range users {
		_, err := repo.CreateUser(user)
		assert.NoError(t, err)
	}

	foundUsers, err := repo.FindByMany("email", "user4@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, foundUsers)
	assert.Len(t, foundUsers, 0)
}
