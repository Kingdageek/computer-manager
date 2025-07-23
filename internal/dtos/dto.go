package dtos

type ComputerDto struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	MacAddress   string  `json:"mac_address"`
	IPAddress    string  `json:"ip_address"`
	EmployeeCode *string `json:"employee_code"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
