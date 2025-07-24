package bootstrap

import (
	"computer-manager/internal/api_clients"
	"computer-manager/internal/config"
)

func InitializeApiClients(cfg *config.Config) *api_clients.ApiClients {
	return api_clients.NewApiClients(cfg.ThirdPartyServices)
}
