package service

import (
	"log"
	"regexp"
)

func IsValidEmail(email, pattern string) bool {
	regex := regexp.MustCompile(pattern)
	log.Println("Validating email:", email, "with pattern:", pattern)
	return regex.MatchString(email)
}
