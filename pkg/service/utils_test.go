package service

import "testing"

func TestIsValidEmail(t *testing.T) {
	emails := []string{
		"test@google.com.ar",
		"cosa@outlook.com",
		"test@asdasd.com",
	}
	for _, email := range emails {
		if !IsValidEmail(email, "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$") {
			t.Errorf("Expected %s to be valid", email)
		}
	}
}

func TestIsInvalidEmail(t *testing.T) {
	emails := []string{
		"test",
		"cosaasdasd",
		"test@",
	}
	for _, email := range emails {
		if IsValidEmail(email, "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$") {
			t.Errorf("Expected %s to be invalid", email)
		}
	}
}
