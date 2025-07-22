package bootstrap

import (
	"computer-manager/internal/config"
	"computer-manager/internal/controllers"
	"computer-manager/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func InitializeControllersAndRoutes(svcs *services.Services, cfg *config.Config) *gin.Engine {
	handlers := controllers.NewControllers(svcs)
	router := InitializeRouter(handlers, cfg)
	log.Println("Controllers and routes initialized successfully")
	return router
}
