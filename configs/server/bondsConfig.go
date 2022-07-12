package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	MONGOCONN string `mapstructure:"MONGODB_CONNSTRING"`
}

type BondsConfigService struct {
	method *BondsConfig
}

type BondsConfig interface {
	LoadConfig(path string) (config Config, err error)
}

func (service *BondsConfigService) LoadConfig(path string) (config Config, err error) {
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
