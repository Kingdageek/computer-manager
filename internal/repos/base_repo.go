package repos

import "gorm.io/gorm"

type Repositories struct {
	db           *gorm.DB
	ComputerRepo *ComputerRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		db:           db,
		ComputerRepo: NewComputerRepository(db),
	}
}

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}
