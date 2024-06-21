package model

import "time"

type Attendance struct {
	AttendanceID int        `db:"attendance_id" json:"attendance_id"`
	EmployeeID   int        `db:"employee_id" json:"employee_id"`
	LocationID   int        `db:"location_id" json:"location_id"`
	AbsentIn     *time.Time `db:"absent_in" json:"absent_in"`
	AbsentOut    *time.Time `db:"absent_out" json:"absent_out"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
	DeletedAt    *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type AttendanceList struct {
	AttendanceID int        `db:"attendance_id" json:"attendance_id"`
	EmployeeID   int        `db:"employee_id" json:"employee_id"`
	LocationID   int        `db:"location_id" json:"location_id"`
	AbsentIn     *time.Time `db:"absent_in" json:"absent_in"`
	AbsentOut    *time.Time `db:"absent_out" json:"absent_out"`
}

type AttendanceReport struct {
	AbsentIn       *time.Time `db:"absent_in" json:"absent_in"`
	AbsentOut      *time.Time `db:"absent_out" json:"absent_out"`
	EmployeeCode   *string    `db:"employee_code" json:"employee_code"`
	EmployeeName   *string    `db:"employee_name" json:"employee_name"`
	DepartmentName *string    `db:"department_name" json:"department_name"`
	PositionName   *string    `db:"position_name" json:"position_name"`
	LocationName   *string    `db:"location_name" json:"location_name"`
}
