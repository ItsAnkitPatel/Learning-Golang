package auth

// won't be accessible outside of the package
func extractSession() string {
	return "LoggedIn"
}

func GetSession() string {

	return extractSession()
}
