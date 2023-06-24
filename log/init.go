package log

import (
	"github.com/veerdone/gopkg/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	sugared *zap.SugaredLogger
)

func Init(conf conf.Log) {
	initSyncFuncMap()
	ws := getSyncs(conf)
	writeSyncer := zapcore.NewMultiWriteSyncer(ws...)

	config := zapcore.EncoderConfig{
		CallerKey:      "caller_line",
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "time",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000"),
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	level := getLevel(conf.Level)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), writeSyncer, zap.NewAtomicLevelAt(level))

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getLevel(level string) zapcore.Level {
	if level == "" {
		return zapcore.InfoLevel
	}
	switch level {
	case "info", "INFO":
		return zapcore.InfoLevel
	case "debug", "DEBUG":
		return zapcore.DebugLevel
	case "warn", "WARN":
		return zapcore.WarnLevel
	case "error", "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Set(l *zap.Logger) {
	logger = l
	sugared = l.Sugar()
}

func GetLogger() *zap.Logger {
	return logger
}

func GetSugared() *zap.SugaredLogger {
	return sugared
}
