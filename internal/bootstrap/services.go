package bootstrap

import (
	"computer-manager/internal/api_clients"
	"computer-manager/internal/config"
	"computer-manager/internal/repos"
	"computer-manager/internal/services"
	"log"
)

func InitializeServices(repos *repos.Repositories, cfg *config.Config, apiClients *api_clients.ApiClients) *services.Services {
	svcs := services.NewServices(repos, apiClients)
	log.Println("Services initialized successfully")
	return svcs
}
