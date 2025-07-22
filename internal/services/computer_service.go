package services

import "computer-manager/internal/repos"

type ComputerService struct {
	repo *repos.ComputerRepository
}

func NewComputerService(repositories *repos.Repositories) *ComputerService {
	return &ComputerService{
		repo: repositories.ComputerRepo,
	}
}
