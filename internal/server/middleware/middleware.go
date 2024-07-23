package middleware

// Initialize middlewares.
func Init(tokenSecret string) {
	initAuth(tokenSecret)
	initLogger()
	initCORS()
}
