package requests

type GetAllComputersRequest struct {
	EmployeeCodes []string `json:"employee_codes"`
}
