package service

import "testing"

func TestIsValidPassword(t *testing.T) {
	passwords := []string{
		"Password123",
		"PassworD1234",
		"PASSWORd123#$%#",
	}
	for _, password := range passwords {
		if !IsValidPassword(password) {
			t.Errorf("Expected %s to be invalid", password)
		}
	}
}

func TestIsInvalidPassword(t *testing.T) {
	passwords := []string{
		"password",
		"PASSWORD",
		"12345678",
		"Password",
	}
	for _, password := range passwords {
		if IsValidPassword(password) {
			t.Errorf("Expected %s to be valid", password)
		}
	}
}
