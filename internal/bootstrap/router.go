package bootstrap

import (
	"computer-manager/internal/api/routes"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	// Create a new Gin router instance
	router := gin.Default()

	// Initialize routes
	routes.CreateRoutes(router)
	return router
}
