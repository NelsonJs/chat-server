package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file:%s \n", err))
	}
	fmt.Println("config app:", viper.GetString("webSocketPort"))
}

func GetViperString(key string) string {
	return viper.GetString(key)
}
