package model

type GenerateJWT struct {
	Token        string `json:"token"`
	ExpiredAt    string `json:"expired_at"`
	EmployeeID   int    `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	EmployeeCode string `json:"employee_code"`
}
