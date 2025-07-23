package repos

import (
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
