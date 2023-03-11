package log

import (
	"github.com/veerdone/gopkg/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	logger  *zap.Logger
	sugared *zap.SugaredLogger
)

func InitLog(conf conf.Log) {
	sync := getSync(conf)
	syncer := zapcore.AddSync(sync)
	writeSyncer := zapcore.NewMultiWriteSyncer(syncer)

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

func getSync(conf conf.Log) io.Writer {
	if conf.Dev || conf.FileName == "" {
		return os.Stdout
	}
	sync := &lumberjack.Logger{
		Filename:   conf.FileName,
		MaxAge:     conf.MaxAge,
		MaxBackups: conf.MaxBackups,
		MaxSize:    conf.MaxSize,
		Compress:   conf.Compress,
	}

	return sync
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
