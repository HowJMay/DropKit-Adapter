package logging

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerLevel int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger interface {
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Debug(msg string, fields ...zapcore.Field)
	Named(name string) Logger
	With(...zapcore.Field) Logger
	Unwrap() *zap.Logger
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(logLevel LoggerLevel, name string, opts ...zap.Option) (Logger, error) {
	var zapLevel zapcore.Level
	switch logLevel {
	case DEBUG:
		zapLevel = zapcore.DebugLevel
	case INFO:
		zapLevel = zapcore.InfoLevel
	case WARN:
		zapLevel = zapcore.WarnLevel
	case ERROR:
		zapLevel = zapcore.ErrorLevel
	case FATAL:
		zapLevel = zapcore.FatalLevel
	default:
		return nil, errors.New("unknown log level")
	}
	level := zap.NewAtomicLevel()
	level.SetLevel(zapLevel)

	// Show the name of caller of the logging function
	opts = append(opts, zap.AddCallerSkip(1))

	config := zap.Config{
		Level:    level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "severity",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "message",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	l, err := config.Build(opts...)
	if err != nil {
		return nil, err
	}

	if name != "" {
		l = l.Named(name)
	}

	return &logger{logger: l}, nil
}

func (l *logger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *logger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *logger) Named(name string) Logger {
	return &logger{logger: l.logger.Named(name)}
}

func (l *logger) With(fields ...zapcore.Field) Logger {
	return &logger{logger: l.logger.With(fields...)}
}

func (l *logger) Unwrap() *zap.Logger {
	return l.logger
}
