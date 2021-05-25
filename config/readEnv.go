package config

import (
	"log"

	"github.com/spf13/viper"
)

func readEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type insertion")
	}

	return value
}

func GetJWTSecret() string {
	return readEnv("JWTSECRET")
}
