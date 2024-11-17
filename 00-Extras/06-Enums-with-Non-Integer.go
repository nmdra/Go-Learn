package main

import "fmt"

type LogLevel string

const (
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARNING"
	Error   LogLevel = "ERROR"
)

func logMessage(level LogLevel, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}

func main() {
	logMessage(Info, "Application started")           // Outputs: [INFO] Application started
	logMessage(Warning, "Low disk space")             // Outputs: [WARNING] Low disk space
	logMessage(Error, "Failed to load configuration") // Outputs: [ERROR] Failed to load configuration
}
