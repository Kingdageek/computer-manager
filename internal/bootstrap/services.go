package bootstrap

import (
	"computer-manager/internal/config"
	"computer-manager/internal/repos"
	"computer-manager/internal/services"
	"log"
)

func InitializeServices(repos *repos.Repositories, cfg *config.Config) *services.Services {
	svcs := services.NewServices(repos)
	log.Println("Services initialized successfully")
	return svcs
}
