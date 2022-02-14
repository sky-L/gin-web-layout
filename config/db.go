package config

type MySqlConfig struct {
	Default DB
}

type DB struct {
	DataSourceName string `mapstructure:"data_source_name" validate:"required,dsn"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}
