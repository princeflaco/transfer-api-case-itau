package infra

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func NewLogger(loggingLevel string) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.TimeKey = "timestamp"

	level := zapcore.InfoLevel
	switch strings.ToLower(loggingLevel) {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(ec), os.Stdout, level)

	logger := zap.New(core, zap.AddCaller())

	return logger
}
