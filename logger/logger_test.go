package logger_test

import (
	"context"
	"testing"

	"github.com/DropKit/DropKit-Adapter/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogger(t *testing.T) {
	type args struct {
		ctx    context.Context
		fields []zapcore.Field
	}

	type test struct {
		name    string
		args    args
		want    []zapcore.Field
		wantErr bool
	}

	tests := []test{
		func() test {
			return test{
				name: "log with empty context",
				args: args{
					ctx: context.Background(),
				},
				want: []zapcore.Field{},
			}
		}(),
		func() test {
			return test{
				name: "log with value context",
				args: args{
					ctx: context.WithValue(context.Background(), "ka", 1),
				},
				want: []zapcore.Field{},
			}
		}(),
		func() test {
			fields := []zapcore.Field{
				zap.String("aa", "bb"),
				zap.String("cc", "dd"),
			}
			return test{
				name: "log with fields only",
				args: args{
					ctx:    context.Background(),
					fields: fields,
				},
				want: fields,
			}
		}(),
		func() test {
			fields := []zapcore.Field{
				zap.String("aa", "bb"),
				zap.String("cc", "dd"),
			}
			ctx, _ := trace.StartSpan(context.Background(), uuid.New().String())

			return test{
				name: "log with span context with extra fields",
				args: args{
					ctx:    ctx,
					fields: fields,
				},
				want: fields,
			}
		}(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fac, logs := observer.New(zapcore.DebugLevel)
			logger, err := logger.NewLogger(logger.DEBUG, zap.WrapCore(func(c zapcore.Core) zapcore.Core { return fac }))

			logger.Info("", tt.args.fields...)

			if tt.wantErr {
				assert.Error(t, err)
				t.Logf("expected error: %+v", err)
				return
			} else if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, tt.want, logs.AllUntimed()[0].Context)
		})
	}
}

func TestLogger_With(t *testing.T) {
	type args struct {
		ctx    context.Context
		fields []zapcore.Field
	}

	type test struct {
		name    string
		args    args
		want    []zapcore.Field
		wantErr bool
	}

	tests := []test{
		func() test {
			return test{
				name: "log with empty context",
				args: args{
					ctx: context.Background(),
				},
				want: []zapcore.Field{},
			}
		}(),
		func() test {
			return test{
				name: "log with value context",
				args: args{
					ctx: context.WithValue(context.Background(), "ka", 1),
				},
				want: []zapcore.Field{},
			}
		}(),
		func() test {
			fields := []zapcore.Field{
				zap.String("aa", "bb"),
				zap.String("cc", "dd"),
			}
			return test{
				name: "log with fields only",
				args: args{
					ctx:    context.Background(),
					fields: fields,
				},
				want: fields,
			}
		}(),
		func() test {
			fields := []zapcore.Field{
				zap.String("aa", "bb"),
				zap.String("cc", "dd"),
			}
			ctx, _ := trace.StartSpan(context.Background(), uuid.New().String())

			return test{
				name: "log with span context with extra fields",
				args: args{
					ctx:    ctx,
					fields: fields,
				},
				want: fields,
			}
		}(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fac, logs := observer.New(zapcore.DebugLevel)
			logger, err := logger.NewLogger(logger.DEBUG, zap.WrapCore(func(c zapcore.Core) zapcore.Core { return fac }))

			logger.With(tt.args.fields...).Info("")

			if tt.wantErr {
				assert.Error(t, err)
				t.Logf("expected error: %+v", err)
				return
			} else if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, tt.want, logs.AllUntimed()[0].Context)
		})
	}
}
