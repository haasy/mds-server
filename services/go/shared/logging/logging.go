package logging

import (
	"github.com/lefinal/meh"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	debugLogger      *zap.Logger
	debugLoggerMutex sync.RWMutex
)

// DebugLogger returns the logger set via SetDebugLogger. If none is set, a
// zap.NewProduction will be created.
func DebugLogger() *zap.Logger {
	debugLoggerMutex.RLock()
	defer debugLoggerMutex.RUnlock()
	if debugLogger == nil {
		tempLogger, _ := zap.NewProduction()
		return tempLogger
	}
	return debugLogger
}

// SetDebugLogger sets the logger that can be retrieved with DebugLogger.
func SetDebugLogger(logger *zap.Logger) {
	debugLoggerMutex.Lock()
	defer debugLoggerMutex.Unlock()
	debugLogger = logger
}

// NewLogger creates a new zap.Logger. Don't forget to call Sync() on the
// returned logged before exiting!
func NewLogger(serviceName string, level zapcore.Level) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.Level = zap.NewAtomicLevelAt(level)
	config.DisableCaller = true
	config.DisableStacktrace = true
	logger, err := config.Build()
	if err != nil {
		return nil, meh.NewInternalErrFromErr(err, "new zap production logger", meh.Details{"config": config})
	}
	return logger, nil
}
