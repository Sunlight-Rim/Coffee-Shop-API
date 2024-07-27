package logger

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

// Create and configure a logger.
func Init(level string, jsonFormat bool) {
	logLevel, err := logger.ParseLevel(level)
	if err != nil {
		logger.Fatalf("parse logger level: %v", err)
	}

	logger.SetLevel(logLevel)
	logger.SetOutput(os.Stdout)

	if jsonFormat {
		logger.SetFormatter(&logger.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05.000",
			DisableHTMLEscape: true,
		})

		return
	}

	logger.SetFormatter(&logger.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		ForceColors:     true,
	})
}
