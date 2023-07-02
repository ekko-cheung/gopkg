/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
