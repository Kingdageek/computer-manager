package repos

import (
	"computer-manager/internal/dtos"
	"computer-manager/internal/models"
	"context"

	"gorm.io/gorm"
)

type ComputerRepository struct {
	BaseRepository
}

func NewComputerRepository(db *gorm.DB) *ComputerRepository {
	return &ComputerRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}

func (cr *ComputerRepository) GetComputerByID(ctx context.Context, id uint, tx *gorm.DB) (*models.Computer, error) {
	dbHandler := cr.PickDbHandler(tx)
	var computer models.Computer
	if err := dbHandler.WithContext(ctx).First(&computer, id).Error; err != nil {
		return nil, err
	}
	return &computer, nil
}

func (cr *ComputerRepository) GetAllComputers(ctx context.Context) ([]*models.Computer, error) {
	var computers []*models.Computer
	if err := cr.db.WithContext(ctx).Find(&computers).Error; err != nil {
		return nil, err
	}
	return computers, nil
}

func (cr *ComputerRepository) DeleteComputerByID(reqCtx context.Context, id uint) error {
	if err := cr.db.WithContext(reqCtx).Delete(&models.Computer{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *ComputerRepository) CountComputersByEmployeeCode(ctx context.Context, employeeCode string) (uint, error) {
	var count int64
	if err := cr.db.WithContext(ctx).Model(&models.Computer{}).Where("employee_code = ?", employeeCode).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}

func (cr *ComputerRepository) dtoToModel(dto *dtos.ComputerDto) *models.Computer {
	return &models.Computer{
		ID:           dto.ID,
		Name:         dto.Name,
		Description:  dto.Description,
		MacAddress:   dto.MacAddress,
		IPAddress:    dto.IPAddress,
		EmployeeCode: dto.EmployeeCode,
	}
}

func (cr *ComputerRepository) CreateComputer(ctx context.Context, dto *dtos.ComputerDto) (*models.Computer, error) {
	computer := cr.dtoToModel(dto)
	if err := cr.db.WithContext(ctx).Create(computer).Error; err != nil {
		return nil, err
	}
	return computer, nil
}

func (cr *ComputerRepository) UpdateComputer(ctx context.Context, id uint, dto *dtos.ComputerDto) (*models.Computer, error) {
	computer := cr.dtoToModel(dto)
	computer.ID = id
	if err := cr.db.WithContext(ctx).Save(computer).Error; err != nil {
		return nil, err
	}
	return computer, nil
}

// MacOrIPAddressExists checks if a computer with the same MAC address or IP address exists, ignoring the computer with the specified ID.
// If idToIgnore is 0, it checks for any existing computer with the same MAC or IP address.
func (cr *ComputerRepository) MacOrIPAddressExists(ctx context.Context, macAddress, ipAddress string, idToIgnore uint) (bool, error) {
	var count int64
	query := cr.db.WithContext(ctx).Model(&models.Computer{}).
		Where("mac_address = ? OR ip_address = ?", macAddress, ipAddress)
	if idToIgnore > 0 {
		query = query.Where("id != ?", idToIgnore)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
