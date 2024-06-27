package config

import (
	"log"

	"github.com/spf13/viper"
)

type Api struct {
	ApiUrl   string `yaml:"ApiUrl"`
	MediaUrl string `yaml:"MediaUrl"`
}

type ServerConfig struct {
	Port int `yaml:"Port"`
}

type Config struct {
	Api    Api          `yaml:"Api"`
	Server ServerConfig `yaml:"Server"`
}

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable decode into struct, %v", err)
	}

	return &cfg
}
