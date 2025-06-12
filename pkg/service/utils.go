package service

import "regexp"

func IsValidEmail(email, pattern string) bool {
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
