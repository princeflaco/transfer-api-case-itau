package util

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	value := ctx.Value("logger")
	if value == nil {
		return nil
	}
	return value.(*zap.Logger)
}

func GetRequestIDFromContext(ctx context.Context) string {
	value := ctx.Value("request_id")
	if value == nil {
		return ""
	}
	return ctx.Value("request_id").(string)
}

func NewTestContext() context.Context {
	logger := NewTestLogger()
	ctx := context.WithValue(context.Background(), "logger", logger)
	ctx = context.WithValue(ctx, "request_id", uuid.NewString())
	return ctx
}
