package app

import (
	"go.uber.org/zap"
)

// Logger .
type Logger struct {
	*zap.SugaredLogger
}

func (c *Context) initLogger() {
	config := zap.NewDevelopmentConfig()
	if c.Config.Logger.Path != "" {
		config.OutputPaths = []string{c.Config.Logger.Path}
	}
	logger, _ := config.Build()

	c.Logger = &Logger{logger.Sugar()}
}

// Print .
func (l *Logger) Print(v ...interface{}) {
	l.Info(v)
}
