package auth

import (
	"hello_bets/intern/model"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := model.User{
		Username: "testuser",
	}
	token, err := GenerateToken(user)
	if err != nil {
		panic(err)
	}
	println("Generated token:", token)
}

func TestValidateToken(t *testing.T) {
	user := model.User{
		Username: "testuser",
	}

	token, err := GenerateToken(user)
	if err != nil {
		panic(err)
	}
	println("Generated token:", token)

	err = ValidateToken(token)
	if err != nil {
		panic(err)
	}

}
