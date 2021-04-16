package config

import (
	"errors"
	"strconv"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
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
