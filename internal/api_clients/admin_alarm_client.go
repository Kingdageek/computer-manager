package api_clients

import "context"

type AdminAlarmClient struct {
	BaseClient
}

func NewAdminAlarmClient(baseURL string) *AdminAlarmClient {
	return &AdminAlarmClient{
		BaseClient: NewBaseClient(baseURL),
	}
}

func (c *AdminAlarmClient) NotifyAdmin(ctx context.Context) {}
