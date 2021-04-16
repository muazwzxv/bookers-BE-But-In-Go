package config

import (
	"errors"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

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

func DBConfig() (DatabaseConfig, error) {
	port, err := strconv.Atoi(readEnv("DBPORT"))
	if err != nil {
		return DatabaseConfig{}, errors.New("Error converting string to int")
	}

	config := DatabaseConfig{
		User:     readEnv("DBUSER"),
		Password: readEnv("DBPASSWORD"),
		Host:     readEnv("DBHOST"),
		DBName:   readEnv("DBNAME"),
		Port:     port,
	}

	return config, nil
}
