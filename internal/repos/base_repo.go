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

// PickDbHandler returns the appropriate database handler based on whether a transaction is provided
func (b *BaseRepository) PickDbHandler(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return b.db
}
