package config

import (
	envConst "KidneySmartBackend/internal/env"
	"encoding/json"
	"errors"

	"os"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Logging struct {
		Level string `json:"level"`
	} `json:"logging"`
	Database struct {
		User    string `json:"user"`
		Name    string `json:"name"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
		SslMode string `json:"sslmode"`
	} `json:"database"`
}

func loadConfig(file string) (Config, error) {
	var config Config

	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
func GetConfig(appEnv string) (Config, error) {
	var configFile string

	switch appEnv {
	case envConst.Prod:
		configFile = "../config/config.prod.json"
	case envConst.Dev:
		configFile = "../config/config.dev.json"
	case envConst.Local:
		configFile = "../config/config.local.json"
	default:
		return Config{}, errors.New("invalid environment: expected Prod, Dev, or Local")
	}

	return loadConfig(configFile)
}
