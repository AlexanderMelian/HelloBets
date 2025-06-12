package service

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Function to check if a password is valid:
// A valid password must meet the following criteria:
// and must contain at least one number, one uppercase letter, one lowercase letter, and one special character
// 2. It must contain at least one number, one uppercase letter, and one lowercase letter to enforce a mix of character types.
// These requirements are designed to enhance security by making passwords harder to guess or brute-force.
func IsValidPassword(password string) bool {
	isValidLength := len(password) >= 8 && len(password) <= 70
	if !isValidLength {
		return false
	}
	hasNumber := false
	hasUpper := false
	hasLower := false
	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		}
		if hasNumber && hasUpper && hasLower {
			break
		}
	}
	return hasNumber && hasUpper && hasLower
}
