package configs

type PostgresConfig struct {
	Host            string `toml:"host" mapstructure:"host"`
	Port            int    `toml:"port" mapstructure:"port"`
	User            string `toml:"user" mapstructure:"user"`
	Password        string `toml:"password" mapstructure:"password"`
	Database        string `toml:"database" mapstructure:"database"`
	MaxOpenConns    int    `toml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns    int    `toml:"max_idle_conns" mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `toml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime int    `toml:"conn_max_idle_time" mapstructure:"conn_max_idle_time"`
}
