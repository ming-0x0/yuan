package configs

type LoggerConfig struct {
	Level string `toml:"level" mapstructure:"level"`
}
