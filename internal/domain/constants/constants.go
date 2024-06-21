package constants

import "errors"

var (
	EmployeeCodeJWTKey = "employee_code"
	EmployeeIDJWTKey   = "employee_id"
	EmployeeNameJWTKey = "employee_name"
)

var (
	ErrService = errors.New("error service: ")
)
