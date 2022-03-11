package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/thinkeridea/go-extend/helper"
	"os"
)

type DB struct {
	DataSourceName string `mapstructure:"data_source_name" validate:"required"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}

type MySqlConfig struct {
	Default DB `mapstructure:"default"`
}

type Config struct {
	MySqlConfig MySqlConfig `mapstructure:"mysql_config"'`
}

func InitConfig() *Config {
	config := &Config{}

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	conf := "./config/conf.yml"
	file := helper.Must(os.Open(conf)).(*os.File)
	defer file.Close()
	helper.Must(nil, viper.ReadConfig(file))

	helper.Must(nil, viper.Unmarshal(config))
	helper.Must(nil, Validate(config))
	return config
}

func Validate(c interface{}) error {
	validate := validator.New()
	validate.SetTagName("validate")
	return validate.Struct(c)
}