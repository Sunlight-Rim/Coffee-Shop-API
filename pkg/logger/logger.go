package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create and configure a logger.
func New(level string, jsonFormat bool) *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Fatalf("parse logger level: %v", err)
	}

	logger.SetLevel(logLevel)

	if jsonFormat {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05.000",
			DisableHTMLEscape: true,
		})

		return logger
	}

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})

	return logger
}
