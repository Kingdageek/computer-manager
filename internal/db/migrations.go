package db

import (
	"computer-manager/internal/models"
	"log"
	"time"

	"gorm.io/gorm"
)

// migrationsRegistry holds a slice of gorm models that need to be migrated.
var migrationsRegistry = []any{
	models.Computer{},
}

type Migration struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);uniqueIndex"`
	AppliedAt time.Time
}

func (Migration) TableName() string {
	return "migrations"
}

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(migrationsRegistry...)

	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
