package viper

import (
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
