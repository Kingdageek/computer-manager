package api_clients

import (
	"context"
	"log"
)

type AdminAlarmClient struct {
	BaseClient
}

func NewAdminAlarmClient(baseURL string) *AdminAlarmClient {
	return &AdminAlarmClient{
		BaseClient: NewBaseClient(baseURL),
	}
}

type NotifyAdminRequest struct {
	Level        string `json:"level"`
	EmployeeAbbr string `json:"employeeAbbreviation"`
	Message      string `json:"message"`
}

func (c *AdminAlarmClient) NotifyAdmin(ctx context.Context, employeeCode, message string) error {
	req := NotifyAdminRequest{
		Level:        "warning",
		EmployeeAbbr: employeeCode,
		Message:      message,
	}

	err := c.Request(ctx, "POST", "/notify", req, nil, nil)
	if err != nil {
		log.Printf("Error notifying admin: %v", err)
		return err
	}
	log.Println("Admin notified successfully")
	return nil
}
