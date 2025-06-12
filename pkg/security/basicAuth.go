package security

func BasicAuth(username, password string) bool {
	if username == "admin" && password == "password" {
		return true
	}
	return false
}
