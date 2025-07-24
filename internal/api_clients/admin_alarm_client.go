package api_clients

import (
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

func (c *AdminAlarmClient) NotifyAdmin(employeeCode, message string) error {
	req := NotifyAdminRequest{
		Level:        "warning",
		EmployeeAbbr: employeeCode,
		Message:      message,
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	err := c.Request("POST", "/notify", req, nil, headers)
	if err != nil {
		log.Printf("Error notifying admin: %v", err)
		return err
	}
	log.Println("Admin notified successfully")
	return nil
}
