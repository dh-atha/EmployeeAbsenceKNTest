package model

import "time"

type Employee struct {
	EmployeeID   int        `db:"employee_id" json:"employee_id"`
	EmployeeCode *string    `db:"employee_code" json:"employee_code"`
	EmployeeName *string    `db:"employee_name" json:"employee_name"`
	Password     string     `db:"password" json:"password"`
	DepartmentID *int       `db:"department_id" json:"department_id"`
	PositionID   *int       `db:"position_id" json:"position_id"`
	Superior     *int       `db:"superior" json:"superior"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type EmployeeList struct {
	EmployeeID   int     `db:"employee_id" json:"employee_id"`
	EmployeeCode *string `db:"employee_code" json:"employee_code"`
	EmployeeName *string `db:"employee_name" json:"employee_name"`
}
