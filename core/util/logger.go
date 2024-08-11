package util

import "go.uber.org/zap"

func NewTestLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
