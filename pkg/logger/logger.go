package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const traceID = "trace_id"

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func (logger *Logger) Info(ctx context.Context, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Info(args...)
}

func (logger *Logger) Infof(ctx context.Context, template string, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Infof(template, args...)
}

func (logger *Logger) Error(ctx context.Context, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Error(args...)
}
func (logger *Logger) Errorf(ctx context.Context, template string, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Errorf(template, args...)
}
func (logger *Logger) Debug(ctx context.Context, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Debug(args...)
}
func (logger *Logger) Debugf(ctx context.Context, template string, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Debugf(template, args...)
}
func (logger *Logger) Warn(ctx context.Context, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Warn(args...)
}
func (logger *Logger) Warnf(ctx context.Context, template string, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Warnf(template, args...)
}

func (logger *Logger) Fatal(ctx context.Context, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Fatal(args...)
}

func (logger *Logger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	logger.With(traceID, getTraceIDFromContext(ctx)).Fatalf(template, args...)
}
func getTraceIDFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value("traceID").(string)
	if !ok {
		return "unknown"
	}
	return traceID
}
