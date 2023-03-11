package log

import "go.uber.org/zap"

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	sugared.Infof(msg, args...)
}

func Debugf(msg string, args ...interface{}) {
	sugared.Debugf(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	sugared.Warnf(msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	sugared.Errorf(msg, args...)
}
