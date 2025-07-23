package services

import (
	"computer-manager/internal/dtos"
	"computer-manager/internal/repos"
	"context"
)

type ComputerService struct {
	repo *repos.ComputerRepository
}

func NewComputerService(repositories *repos.Repositories) *ComputerService {
	return &ComputerService{
		repo: repositories.ComputerRepo,
	}
}

func (cs *ComputerService) GetComputerByID(ctx context.Context, id uint) (*dtos.ComputerDto, error) {
	computer, err := cs.repo.GetComputerByID(ctx, id, nil)
	if err != nil {
		return nil, err
	}
	return computer.ToDto(), nil
}

func (cs *ComputerService) GetAllComputers(ctx context.Context) ([]*dtos.ComputerDto, error) {
	computers, err := cs.repo.GetAllComputers(ctx)
	if err != nil {
		return nil, err
	}

	computerDtos := make([]*dtos.ComputerDto, len(computers))
	for i, computer := range computers {
		computerDtos[i] = computer.ToDto()
	}
	return computerDtos, nil
}
