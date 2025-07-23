package bootstrap

import (
	"computer-manager/internal/api/routes"
	"computer-manager/internal/config"
	"computer-manager/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(handlers *controllers.Controllers, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	routes.CreateRoutes(router, handlers)
	return router
}
