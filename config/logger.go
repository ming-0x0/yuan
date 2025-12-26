package config

type LoggerConfig struct {
	Level string `toml:"level" mapstructure:"level"`
}

type BunConfig struct {
	SlowThreshold     int    `toml:"slow_threshold" mapstructure:"slow_threshold"`
	IgnoreNoRowsError bool   `toml:"ignore_no_rows_error" mapstructure:"ignore_no_rows_error"`
	Level             string `toml:"level" mapstructure:"level"`
}
