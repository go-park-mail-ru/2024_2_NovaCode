package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMapping = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.WarnLevel,
	"error": zap.ErrorLevel,
}

type Logger interface {
	Debug(msg string, f ...zap.Field)
	Info(msg string, f ...zap.Field)
	Warn(msg string, f ...zap.Field)
	Error(msg string, f ...zap.Field)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type Log struct {
	ctx context.Context
	log *zap.Logger
}

type loggerKey struct{}

var encoderCfg = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func New(cfg *config.LoggerConfig) Logger {
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.Lock(os.Stdout), zap.DebugLevel)
	if cfg.Format == "json" {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.Lock(os.Stdout), zap.DebugLevel)
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	if v, ok := levelMapping[cfg.Level]; ok {
		zap.NewAtomicLevelAt(v)
	} else {
		zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	ctx := context.WithValue(context.Background(), loggerKey{}, logger)
	return &Log{ctx, logger}
}

func ctxLogger(ctx context.Context) *zap.Logger {
	if logger, _ := ctx.Value(loggerKey{}).(*zap.Logger); logger != nil {
		return logger
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		zap.DebugLevel,
	), zap.AddCaller(), zap.AddCallerSkip(1))
	zap.NewAtomicLevelAt(zap.InfoLevel)

	return logger
}

func (l *Log) Debug(msg string, f ...zap.Field) {
	ctxLogger(l.ctx).Debug(msg, f...)
}

func (l *Log) Info(msg string, f ...zap.Field) {
	ctxLogger(l.ctx).Info(msg, f...)
}

func (l *Log) Warn(msg string, f ...zap.Field) {
	ctxLogger(l.ctx).Warn(msg, f...)
}

func (l *Log) Error(msg string, f ...zap.Field) {
	ctxLogger(l.ctx).Error(msg, f...)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	ctxLogger(l.ctx).Debug(msg)
}

func (l *Log) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	ctxLogger(l.ctx).Info(msg)
}

func (l *Log) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	ctxLogger(l.ctx).Warn(msg)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	ctxLogger(l.ctx).Error(msg)
}
