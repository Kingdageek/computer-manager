package main

import (
	"computer-manager/internal/bootstrap"
)

// @title Computer Manager API
// @version 1.0
// @description API for sysadmins to manage computers in an organization.
// @host localhost:8000
// @BasePath /api/v1
func main() {
	bootstrap.LoadEnv()
	cfg := bootstrap.InitializeConfig()
	db := bootstrap.InitializeDB(cfg)
	repos := bootstrap.InitializeRepos(db)
	apiClients := bootstrap.InitializeApiClients(cfg)
	services := bootstrap.InitializeServices(repos, cfg, apiClients)
	router := bootstrap.InitializeControllersAndRoutes(services, cfg)
	router.Run()
}
