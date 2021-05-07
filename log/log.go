package log

import (
	"go.uber.org/zap"
)

var sugaredLogger *zap.SugaredLogger

func init() {
	sugaredLogger = logger.Sugar()
}

func Debug(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}
