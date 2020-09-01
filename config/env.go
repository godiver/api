package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Env() {
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$App")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
