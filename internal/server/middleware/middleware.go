package middleware

// Initialize middlewares.
func Init(tokenSecret, employeeToken string) {
	initAuth(tokenSecret, employeeToken)
	initLogger()
	initCORS()
}
