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
	logger, err := logger.New(&cfg.Log)
	if err != nil {
		return nil, fmt.Errorf("starting logger: %w", err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)

	return &App{
		Log: logger,
	}, nil
}
