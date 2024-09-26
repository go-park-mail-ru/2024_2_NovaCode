package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
)

var levelMapping = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

type Logger interface {
	Debug(msg string, f ...slog.Attr)
	Info(msg string, f ...slog.Attr)
	Warn(msg string, f ...slog.Attr)
	Error(msg string, f ...slog.Attr)
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type Log struct {
	ctx  context.Context
	slog *slog.Logger
}

func New(cfg *config.LoggerConfig) Logger {
	options := &slog.HandlerOptions{}

	if v, ok := levelMapping[cfg.Level]; ok {
		options.Level = v
	} else {
		options.Level = slog.LevelInfo
	}

	var handler slog.Handler

	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, options)
	default:
		handler = slog.NewTextHandler(os.Stderr, options)
	}

	log := slog.New(handler)

	return &Log{context.Background(), log}
}

func (l *Log) Debug(msg string, f ...slog.Attr) {
	l.slog.LogAttrs(l.ctx, slog.LevelDebug, msg, f...)
}

func (l *Log) Info(msg string, f ...slog.Attr) {
	l.slog.LogAttrs(l.ctx, slog.LevelInfo, msg, f...)
}

func (l *Log) Warn(msg string, f ...slog.Attr) {
	l.slog.LogAttrs(l.ctx, slog.LevelWarn, msg, f...)
}

func (l *Log) Error(msg string, f ...slog.Attr) {
	l.slog.LogAttrs(l.ctx, slog.LevelError, msg, f...)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.slog.LogAttrs(l.ctx, slog.LevelDebug, msg)
}

func (l *Log) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.slog.LogAttrs(l.ctx, slog.LevelInfo, msg)
}

func (l *Log) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.slog.LogAttrs(l.ctx, slog.LevelWarn, msg)
}

func (l *Log) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.slog.LogAttrs(l.ctx, slog.LevelError, msg)
}
