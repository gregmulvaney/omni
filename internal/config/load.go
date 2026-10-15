package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load() {
}

func RegisterFlags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("omni", pflag.ContinueOnError)

	fs.String("log.level", "info", "Set log level (debug, info, warn, error)")
	fs.String("log.format", "json", "Set log format (json, console)")
	fs.Bool("log.dev", false, "Enable development logging")

	return fs
}

func WriteDefault(path string, force bool) error {
	v := viper.New()

	if err := v.BindPFlags(RegisterFlags()); err != nil {
		return fmt.Errorf("binding flags: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	write := v.SafeWriteConfigAs
	if force {
		write = v.WriteConfigAs
	}

	if err := write(path); err != nil {
		return fmt.Errorf("writing config file: %w", err)
	}

	return nil
}
