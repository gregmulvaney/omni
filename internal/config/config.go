package config

const (
	DefaultPath = "./config/omni.yaml"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server" yaml:"server"`
	Database DatabaseConfig `mapstructure:"db"`
	Log      LogConfig      `mapstructure:"log" yaml:"log"`
}

type ServerConfig struct {
	Host       string `mapstructure:"host" yaml:"host"`
	Port       int    `mapstructure:"port" yaml:"port"`
	SecurePort int    `mapstructure:"secure-port" yaml:"secure-port"`
}

type DatabaseConfig struct{}

type LogConfig struct {
	Format      string `mapstructure:"format" yaml:"format"`
	Level       string `mapstructure:"level" yaml:"level"`
	Development bool   `mapstructure:"dev" yaml:"dev"`
}
