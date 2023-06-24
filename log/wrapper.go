package log

import (
	"context"
	"github.com/veerdone/gopkg"
	"go.uber.org/zap"
)

type Logger struct {
	l *zap.Logger
	s *zap.SugaredLogger
}

func NewLogger(l *zap.Logger, s *zap.SugaredLogger) *Logger {
	return &Logger{l: l, s: s}
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
	value := ctx.Value(gopkg.TraceId)
	if value != nil {
		return l.l.With(zap.Any(gopkg.TraceId, value))
	}

	return l.l
}

func (l *Logger) Infof(tmp string, args ...interface{}) {
	l.s.Infof(tmp, args...)
}

func (l *Logger) Warnf(tmp string, args ...interface{}) {
	l.s.Warnf(tmp, args...)
}

func (l *Logger) Errorf(tmp string, args ...interface{}) {
	l.s.Errorf(tmp, args...)
}

func (l *Logger) Fatalf(tmp string, args ...interface{}) {
	l.s.Fatalf(tmp, args...)
}

func (l *Logger) Debugf(tmp string, args ...interface{}) {
	l.s.Debugf(tmp, args...)
}