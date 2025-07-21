package bootstrap

import (
	"computer-manager/internal/config"
	"computer-manager/internal/db"
	"log"

	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dbInstance, err := db.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatal("Failed to initialize database",
			err,
			"host: ", cfg.Database.Host,
			"port: ", cfg.Database.Port,
			"database", cfg.Database.Name,
		)
	}
	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatal("Failed to get database instance", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database", err)
	}
	log.Println("Database connection established successfully")
	return dbInstance
}
