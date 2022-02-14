package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/thinkeridea/go-extend/helper"
	"os"
)

type Config struct {
	MySqlConfig MySqlConfig
}

func InitConfig() *Config {
	config := Config{}

	viper.AutomaticEnv()
	// file config
	viper.SetConfigType("yml")

	conf := "./config/conf.yml"
	// parse config
	file := helper.Must(os.Open(conf)).(*os.File)
	defer file.Close()
	helper.Must(nil, viper.ReadConfig(file))
	helper.Must(nil, viper.Unmarshal(config))

	fmt.Println(config)

	return &config
}
