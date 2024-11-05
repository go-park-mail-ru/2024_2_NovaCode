package logger

import (
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/google/uuid"
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
	Debug(msg string, requestID any, f ...zap.Field)
	Info(msg string, requestID any, f ...zap.Field)
	Warn(msg string, requestID any, f ...zap.Field)
	Error(msg string, requestID any, f ...zap.Field)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type Log struct {
	log *zap.Logger
}

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
	EncodeCaller:   zapcore.FullCallerEncoder,
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

	return &Log{logger}
}

func (l *Log) Debug(msg string, requestIDAny any, f ...zap.Field) {
	requestID, ok := requestIDAny.(uuid.UUID)
	if !ok {
		l.log.Debug(msg, f...)
		return
	}
	f = append(f, zap.String("request_id", requestID.String()))
	l.log.Debug(msg, f...)
}

func (l *Log) Info(msg string, requestIDAny any, f ...zap.Field) {
	requestID, ok := requestIDAny.(uuid.UUID)
	if !ok {
		l.log.Info(msg, f...)
		return
	}
	f = append(f, zap.String("request_id", requestID.String()))
	l.log.Info(msg, f...)
}

func (l *Log) Warn(msg string, requestIDAny any, f ...zap.Field) {
	requestID, ok := requestIDAny.(uuid.UUID)
	if !ok {
		l.log.Warn(msg, f...)
		return
	}
	f = append(f, zap.String("request_id", requestID.String()))
	l.log.Warn(msg, f...)
}

func (l *Log) Error(msg string, requestIDAny any, f ...zap.Field) {
	requestID, ok := requestIDAny.(uuid.UUID)
	if !ok {
		l.log.Error(msg, f...)
		return
	}
	f = append(f, zap.String("request_id", requestID.String()))
	l.log.Error(msg, f...)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.log.Debug(msg)
}

func (l *Log) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.log.Info(msg)
}

func (l *Log) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.log.Warn(msg)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.log.Error(msg)
}
