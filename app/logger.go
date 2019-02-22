package app

import (
	"go.uber.org/zap"
)

// Logger .
type Logger struct {
	*zap.SugaredLogger
}

// Print .
func (l *Logger) Print(v ...interface{}) {
	l.Info(v)
}
