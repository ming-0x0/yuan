package configs

import (
	"strings"

	"github.com/spf13/viper"
)

type Env string

const (
	Local Env = "local"
	Dev   Env = "dev"
	Stg   Env = "stg"
	Prod  Env = "prod"
)

func (e Env) IsProd() bool {
	return e == Prod
}

type Config struct {
	Env      Env             `toml:"env" mapstructure:"env"`
	Logger   *LoggerConfig   `toml:"logger" mapstructure:"logger"`
	Postgres *PostgresConfig `toml:"postgres" mapstructure:"postgres"`
}

func LoadConfig() (*Config, error) {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("configs")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Enable environment variables to override config
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
