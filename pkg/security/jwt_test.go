package security

import (
	"hello_bets/pkg/model"
	"testing"

	"github.com/golang-jwt/jwt"
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
	//println("Generated token:", token)
	err = ValidateToken(token)
	if err != nil {
		panic(err)
	}
}

func TestValidateTokenExpired(t *testing.T) {
	user := model.User{
		Username: "testuser",
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      0,
		"iat":      0,
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenObj.SignedString(secretKey)
	if err != nil {
		panic(err)
	}
	err = ValidateToken(tokenString)
	if err == nil {
		t.Error("Expected token validation to fail for expired token, but it succeeded")
	} else {
		println("Expected error for expired token:", err.Error())
	}
}
