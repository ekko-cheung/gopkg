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
