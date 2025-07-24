package services

import (
	"computer-manager/internal/api_clients"
	"computer-manager/internal/repos"
)

type Services struct {
	ComputerService *ComputerService
}

func NewServices(repos *repos.Repositories, apiClients *api_clients.ApiClients) *Services {
	computerService := NewComputerService(repos, apiClients)
	return &Services{
		ComputerService: computerService,
	}
}
