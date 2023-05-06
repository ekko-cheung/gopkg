package log

import (
	"context"
	"github.com/veerdone/gopkg"
	"go.uber.org/zap"
)

func WithContext(ctx context.Context) *zap.Logger {
	val := ctx.Value(gopkg.TraceId)
	if val != nil {
		return logger.With(zap.Any(gopkg.TraceId, val))
	}

	return logger
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func CtxInfo(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func CtxWarn(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func CtxDebug(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Debug(msg, fields...)
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
