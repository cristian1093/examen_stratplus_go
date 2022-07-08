package tools

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string) (value string) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error with viper ", err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return
}

// GetEnvString get enviroment string values
func GetEnvString(key string) (value string) {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error in config file: %s \n", err))
		return
	}

	return viper.GetString(key)
}
