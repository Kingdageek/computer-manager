package routes

import (
	"computer-manager/internal/api/controllers"

	_ "computer-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRoutes(router *gin.Engine, handlers *controllers.Controllers) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")

	v1.GET("/computers", handlers.Computer.GetAll)
	v1.GET("/computers/:id", handlers.Computer.GetByID)
	v1.POST("/computers", handlers.Computer.Create)
	v1.PUT("/computers/:id", handlers.Computer.Update)
	v1.DELETE("/computers/:id", handlers.Computer.Delete)
}
