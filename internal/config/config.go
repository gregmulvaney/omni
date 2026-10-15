package config

const (
	DefaultPath = "./config/omni.yaml"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"db"`
	Log      LogConfig      `mapstructure:"log" yaml:"log"`
}

type ServerConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	SecurePort int    `mapstructure:"secure-port"`
}

type DatabaseConfig struct{}

type LogConfig struct {
	Format      string `mapstructure:"format" yaml:"format"`
	Level       string `mapstructure:"level" yaml:"level"`
	Development bool   `mapstructure:"dev" yaml:"dev"`
}
