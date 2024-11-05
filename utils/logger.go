package utils

import (
	"log"
	"os"
)

// InitLogger initializes the logger with a specified log file.
func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

// Info logs informational messages.
func Info(msg string) {
	log.Println("INFO:", msg)
}

// Error logs error messages.
func Error(msg string) {
	log.Println("ERROR:", msg)
}

// Fatal logs fatal error messages and exits the application.
func Fatal(msg string) {
	log.Fatalln("FATAL:", msg)
}
