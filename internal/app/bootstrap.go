package app

import (
	"fmt"

	"github.com/gregmulvaney/omni/internal/config"
	"github.com/gregmulvaney/omni/internal/logger"
	"go.uber.org/zap"
)

type App struct {
	Log *zap.Logger
}

func Bootstrap(cfg *config.Config) (*App, error) {
	// Create logger instance
	log, err := logger.New(&cfg.Log)
	if err != nil {
		return nil, fmt.Errorf("starting logger: %w", err)
	}
	zap.ReplaceGlobals(log)
	zap.RedirectStdLog(log)

	return &App{
		Log: log,
	}, nil
}

// Close flushes buffered log entries. Call it once at shutdown.
func (a *App) Close() {
	// Sync commonly fails on stderr/stdout (e.g. "inappropriate ioctl"); not actionable.
	_ = a.Log.Sync()
}
