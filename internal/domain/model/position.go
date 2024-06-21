package model

import "time"

type Position struct {
	PositionID   int        `db:"position_id" json:"position_id"`
	DepartmentID *int       `db:"department_id" json:"department_id"`
	PositionName *string    `db:"position_name" json:"position_name"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
	DeletedAt    *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type PositionList struct {
	PositionID   int     `db:"position_id" json:"position_id"`
	DepartmentID *int    `db:"department_id" json:"department_id"`
	PositionName *string `db:"position_name" json:"position_name"`
}
