package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type LoggerLevel int

func NewLogger(logLevel LoggerLevel, name string, opts ...zap.Option) (*zap.Logger, error) {
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

	return l, nil
}
