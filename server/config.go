package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MONGOCONN string `mapstructure:"MONGODB_CONNSTRING"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err in config.go: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
