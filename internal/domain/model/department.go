package model

import "time"

type Department struct {
	DepartmentID   int        `db:"department_id" json:"department_id"`
	DepartmentName *string    `db:"department_name" json:"department_name"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	CreatedBy      *string    `db:"created_by" json:"created_by"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy      *string    `db:"updated_by" json:"updated_by"`
	DeletedAt      *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type DepartmentList struct {
	DepartmentID   int     `db:"department_id" json:"department_id"`
	DepartmentName *string `db:"department_name" json:"department_name"`
}
