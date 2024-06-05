package logger

import "log"

// Logger provides logging functions.
type Logger struct {
}

// NewLogger creates a new Logger.
func NewLogger() *Logger {
	return &Logger{}
}

// Info logs informational messages.
func (l *Logger) Info(args ...interface{}) {
	log.Println(args...)
}

// Error logs error messages.
func (l *Logger) Error(args ...interface{}) {
	log.Println(args...)
}
