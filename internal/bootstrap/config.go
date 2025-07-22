package bootstrap

import (
	"computer-manager/internal/config"
	"log"
)

func InitializeConfig() *config.Config {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error initializing config", err)
	}
	return config
}
