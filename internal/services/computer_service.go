package services

import (
	"computer-manager/internal/api/http_errors"
	"computer-manager/internal/api_clients"
	"computer-manager/internal/dtos"
	"computer-manager/internal/helpers"
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
	exists, err := cs.repo.MacOrIPAddressExists(ctx, dto.MacAddress, dto.IPAddress, 0)
	if err != nil {
		log.Printf("Error checking MAC/IP existence: %v", err)
		return nil, err
	}
	if exists {
		return nil, http_errors.NewBadRequestError("Computer with the same MAC or IP address already exists")
	}
	computer, err := cs.repo.CreateComputer(ctx, dto)
	if err != nil {
		log.Printf("Error creating computer: %v", err)
		return nil, err
	}
	if dto.EmployeeCode != nil && *dto.EmployeeCode != "" {
		go cs.runEmployeeAssignmentCheck(ctx, *dto.EmployeeCode)
	}
	return computer.ToDto(), nil
}

func (cs *ComputerService) DeleteComputer(reqCtx context.Context, id uint) (bool, error) {
	err := cs.repo.DeleteComputerByID(reqCtx, id)
	if err != nil {
		log.Printf("Error deleting computer with ID %d: %v", id, err)
		return false, err
	}
	return true, nil
}

func (cs *ComputerService) UpdateComputer(ctx context.Context, id uint, dto *dtos.ComputerDto) (*dtos.ComputerDto, error) {
	// confirms that the computer exists
	prevComputer, err := cs.GetComputerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	exists, err := cs.repo.MacOrIPAddressExists(ctx, dto.MacAddress, dto.IPAddress, id)
	if err != nil {
		log.Printf("Error checking MAC/IP existence: %v", err)
		return nil, err
	}
	if exists {
		return nil, http_errors.NewBadRequestError("Computer with the same MAC or IP address already exists")
	}
	computer, err := cs.repo.UpdateComputer(ctx, id, dto)
	if err != nil {
		log.Printf("Error updating computer with ID %d: %v", id, err)
		return nil, err
	}

	// Check if the employee code has changed and it's different from what was previously assigned
	if dto.EmployeeCode != nil && *dto.EmployeeCode != "" {
		if prevComputer.EmployeeCode == nil || (prevComputer != nil && *prevComputer.EmployeeCode != *dto.EmployeeCode) {
			go cs.runEmployeeAssignmentCheck(ctx, *dto.EmployeeCode)
		}
	}
	return computer.ToDto(), nil
}

func (cs *ComputerService) runEmployeeAssignmentCheck(ctx context.Context, employeeCode string) error {
	count, err := cs.repo.CountComputersByEmployeeCode(ctx, employeeCode)
	if err != nil {
		log.Printf("Error counting computers for employee code %s: %v", employeeCode, err)
		return err
	}
	adminAlertThreshold := helpers.GetEnv("ADMIN_ALERT_THRESHOLD", 3)
	if count >= uint(adminAlertThreshold) {
		log.Printf("Employee code %s has %d computers assigned, exceeding the threshold of %d", employeeCode, count, adminAlertThreshold)
		message := fmt.Sprintf("Employee code %s has too many computers assigned (%d)", employeeCode, count)
		err := cs.adminAlarmClient.NotifyAdmin(ctx, employeeCode, message)
		if err != nil {
			log.Printf("Error sending alarm for employee code %s: %v", employeeCode, err)
			return err
		}
	}
	return nil
}
