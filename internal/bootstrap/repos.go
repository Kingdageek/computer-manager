package bootstrap

import (
	"computer-manager/internal/repos"
	"log"

	"gorm.io/gorm"
)

func InitializeRepos(db *gorm.DB) *repos.Repositories {
	repositories := repos.NewRepositories(db)
	log.Println("Repositories initialized successfully")
	return repositories
}
