package model

import "time"

type Location struct {
	LocationID   int        `db:"location_id" json:"location_id"`
	LocationName *string    `db:"location_name" json:"location_name"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	CreatedBy    *string    `db:"created_by" json:"created_by"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    *string    `db:"updated_by" json:"updated_by"`
	DeletedAt    *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type LocationList struct {
	LocationID   int     `db:"location_id" json:"location_id"`
	LocationName *string `db:"location_name" json:"location_name"`
}
