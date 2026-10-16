package logger

import (
	"fmt"
	"strings"

	"github.com/gregmulvaney/omni/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *config.LogConfig) (*zap.Logger, error) {
	var zc zap.Config

	// Set log format
	switch strings.ToLower(cfg.Format) {
	case "json":
		zc = zap.NewProductionConfig()
	case "console":
		zc = zap.NewDevelopmentConfig()
		zc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		return nil, fmt.Errorf("invalid log format: %s", cfg.Format)
	}

	level, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}
	zc.Level = zap.NewAtomicLevelAt(level)

	return zc.Build()
}

func parseLevel(s string) (zapcore.Level, error) {
	var level zapcore.Level
	if err := level.Set(strings.ToLower(s)); err != nil {
		return zapcore.InfoLevel, fmt.Errorf("invalid log level %q: %w", s, err)
	}
	return level, nil
}
