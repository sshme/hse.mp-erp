package utils

import "log"

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type StdLogger struct {
	*log.Logger
}

// NewLogger — StdLogger как Logger.
func NewLogger() Logger {
	return &StdLogger{log.Default()}
}
