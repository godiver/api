package env

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config has all environment variables
type env struct {
	API API
}

var Env = env{}

func ReadEnv() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$App")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&Env)
}
