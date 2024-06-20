package middleware

import "github.com/sirupsen/logrus"

// Initialize middlewares.
func Init(logger logrus.FieldLogger, tokenSecret string) {
	initAuth(tokenSecret)
	initLogger(logger)
	initCORS()
}
