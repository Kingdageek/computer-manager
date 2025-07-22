package services

import "computer-manager/internal/repos"

type Services struct {
	ComputerService *ComputerService
}

func NewServices(repos *repos.Repositories) *Services {
	computerService := NewComputerService(repos)
	return &Services{
		ComputerService: computerService,
	}
}
