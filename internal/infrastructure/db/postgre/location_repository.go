package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type LocationRepositoryImpl struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) repository.LocationRepository {
	return &LocationRepositoryImpl{db: db}
}

func (r *LocationRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Location, error) {
	query := "SELECT * FROM Locations WHERE location_id = $1 and deleted_at is null"

	var emp model.Location
	err := r.db.GetContext(ctx, &emp, query, id)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *LocationRepositoryImpl) GetAll(ctx context.Context) ([]*model.LocationList, error) {
	var Locations []*model.LocationList
	query := "SELECT location_id, location_name FROM locations WHERE deleted_at IS NULL"

	err := r.db.SelectContext(ctx, &Locations, query)
	if err != nil {
		return nil, err
	}

	return Locations, nil
}

func (r *LocationRepositoryImpl) Create(ctx context.Context, emp *model.Location) (int, error) {
	query := "INSERT INTO locations (location_name, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5) RETURNING Location_id"

	var LocationID int
	err := r.db.QueryRowContext(ctx, query,
		emp.LocationName,
		emp.CreatedAt,
		emp.CreatedBy,
		emp.UpdatedAt,
		emp.UpdatedBy,
	).Scan(&LocationID)
	if err != nil {
		return 0, err
	}

	return LocationID, nil
}

func (r *LocationRepositoryImpl) Update(ctx context.Context, emp *model.Location) error {
	query := `
		UPDATE Locations SET
			location_name = COALESCE(NULLIF(:location_name, ''), location_name),
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE location_id = :location_id AND deleted_at IS NULL
	`

	res, err := r.db.NamedExecContext(ctx, query, emp)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}

func (r *LocationRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "UPDATE locations SET deleted_at = $1 WHERE location_id = $2 and deleted_at is null"

	res, err := r.db.ExecContext(ctx, query, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}
