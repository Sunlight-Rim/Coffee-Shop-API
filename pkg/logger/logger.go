package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create and configure a logger.
func Init(level string, jsonFormat bool) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatalf("parse logger level: %v", err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetOutput(os.Stdout)

	if jsonFormat {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05.000",
			DisableHTMLEscape: true,
		})

		return
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		ForceColors:     true,
	})
}
