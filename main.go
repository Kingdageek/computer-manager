package main

import (
	"computer-manager/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.LoadEnv()
	cfg := bootstrap.InitializeConfig()
	db := bootstrap.InitializeDB(cfg)
	repos := bootstrap.InitializeRepos(db)
	apiClients := bootstrap.InitializeApiClients(cfg)
	services := bootstrap.InitializeServices(repos, cfg, apiClients)
	router := bootstrap.InitializeControllersAndRoutes(services, cfg)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
