package accessory

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

// NewLogger creates a New Logger.
func NewLogger(traceID, spanID string) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"stdout",
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

// Log implements a common logger pattern in golang.
func Log(ctx context.Context) *zap.Logger {
	return ctx.Value("logger").(*zap.Logger)
}
