package repos

import "gorm.io/gorm"

type ComputerRepository struct {
	BaseRepository
}

func NewComputerRepository(db *gorm.DB) *ComputerRepository {
	return &ComputerRepository{
		BaseRepository: *NewBaseRepository(db),
	}
}
