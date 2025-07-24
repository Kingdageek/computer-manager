package services

import (
	"computer-manager/internal/api/http_errors"
	"computer-manager/internal/api_clients"
	"computer-manager/internal/dtos"
	"computer-manager/internal/repos"
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ComputerService struct {
	repo             *repos.ComputerRepository
	adminAlarmClient *api_clients.AdminAlarmClient
}

func NewComputerService(repositories *repos.Repositories, apiClients *api_clients.ApiClients) *ComputerService {
	return &ComputerService{
		repo:             repositories.ComputerRepo,
		adminAlarmClient: apiClients.AdminAlarm,
	}
}

func (cs *ComputerService) GetComputerByID(ctx context.Context, id uint) (*dtos.ComputerDto, error) {
	computer, err := cs.repo.GetComputerByID(ctx, id, nil)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http_errors.NewNotFoundError(fmt.Sprintf("Computer with ID %d not found", id))
		}
		log.Printf("Error fetching computer by ID %d: %v", id, err)
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

func (cs *ComputerService) CreateComputer(ctx context.Context, dto *dtos.ComputerDto) (*dtos.ComputerDto, error) {
	panic("unimplemented")
}

func (cs *ComputerService) DeleteComputer(reqCtx context.Context, id uint) (bool, error) {
	err := cs.repo.DeleteComputerByID(reqCtx, id)
	if err != nil {
		log.Printf("Error deleting computer with ID %d: %v", id, err)
		return false, err
	}
	return true, nil
}

func (cs *ComputerService) UpdateComputer(ctx context.Context, computerId uint, dto *dtos.ComputerDto) (*dtos.ComputerDto, error) {
	panic("unimplemented")
}
