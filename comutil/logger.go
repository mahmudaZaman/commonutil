package comutil

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
)

type contextKeyType string

// Logger is the key for holding zap logger in the context.
const Logger contextKeyType = "logger"

// NewLogger creates a New Logger.
func NewLogger(traceID, spanID string) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"stdout",
		"account.log",
	}
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = fmt.Sprintf("message[%s]", traceID)
	l, _ := config.Build()
	return l.With(zap.String("traceID", traceID),
		zap.String("spanID", spanID),
	)
}

// NewTraceableLogger accepts the input of c.Get("logger") and returns zap logger.
// To avoid gin context in the common I think this is better approach to extract logger from gin context.
func NewTraceableLogger(ins interface{}, exists bool) *zap.Logger {
	if !exists {
		log.Fatal("Application configuration error, zap.Logger is not present in gin context")
	}
	logger, ok := ins.(*zap.Logger)
	if !ok {
		log.Fatal("Application configuration error, invalid zap.Logger is set as logger in gin context")
	}
	return logger
}

// Log implements a common logger pattern in golang.
func Log(ctx context.Context) *zap.Logger {
	return ctx.Value(Logger).(*zap.Logger)
}
