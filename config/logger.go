package config

type LoggerConfig struct {
	Level string `toml:"level" mapstructure:"level"`
}
