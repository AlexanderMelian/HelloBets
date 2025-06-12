package security

import (
	"hello_bets/pkg/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte

func init() {
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	if secretKey == nil {
		panic("secretKey is nil")
	}
}

func GenerateToken(u model.User) (string, error) {
	claims := jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 4).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
			}
		}
		return nil
	}
	return jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}
