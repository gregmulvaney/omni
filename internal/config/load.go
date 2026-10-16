package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load(v *viper.Viper, fs *pflag.FlagSet, file string) (*Config, error) {
	// Bind pflags from viper instance
	if err := v.BindPFlags(fs); err != nil {
		return nil, fmt.Errorf("bindng flags: %w", err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("omni")

	v.AutomaticEnv()

	// TODO: read in config file

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}

	// TODO: Normalize relevant values

	// TODO: Implement validation for all values

	return &cfg, nil
}

func RegisterFlags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("omni", pflag.ContinueOnError)

	// Log settings
	fs.String("log.level", "info", "Set log level (debug, info, warn, error)")
	fs.String("log.format", "json", "Set log format (json, console)")
	fs.Bool("log.dev", false, "Enable development logging")

	// Server settings
	fs.String("server.host", "0.0.0.0", "Server host address")
	fs.Int("server.port", 6767, "Server cleartext port")
	fs.Int("server.secure-port", 4567, "Server secure port")

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
