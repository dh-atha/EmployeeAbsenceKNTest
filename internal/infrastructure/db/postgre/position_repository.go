package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type PositionRepositoryImpl struct {
	db *sqlx.DB
}

func NewPositionRepository(db *sqlx.DB) repository.PositionRepository {
	return &PositionRepositoryImpl{db: db}
}

func (r *PositionRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Position, error) {
	query := "SELECT * FROM positions WHERE position_id = $1 and deleted_at is null"

	var emp model.Position
	err := r.db.GetContext(ctx, &emp, query, id)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *PositionRepositoryImpl) GetAll(ctx context.Context) ([]*model.PositionList, error) {
	var Positions []*model.PositionList
	query := "SELECT position_id, department_id, position_name FROM positions WHERE deleted_at IS NULL"

	err := r.db.SelectContext(ctx, &Positions, query)
	if err != nil {
		return nil, err
	}

	return Positions, nil
}

func (r *PositionRepositoryImpl) Create(ctx context.Context, emp *model.Position) (int, error) {
	fmt.Println("hi")

	query := "INSERT INTO positions (position_name, department_id, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING position_id"

	var PositionID int
	err := r.db.QueryRowContext(ctx, query,
		emp.PositionName,
		emp.DepartmentID,
		emp.CreatedAt,
		emp.CreatedBy,
		emp.UpdatedAt,
		emp.UpdatedBy,
	).Scan(&PositionID)
	if err != nil {
		return 0, err
	}

	return PositionID, nil
}

func (r *PositionRepositoryImpl) Update(ctx context.Context, emp *model.Position) error {
	query := `
		UPDATE positions SET
			position_name = COALESCE(NULLIF(:position_name, ''), position_name),
			department_id = COALESCE(NULLIF(:department_id, 0), department_id),
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE position_id = :position_id AND deleted_at IS NULL
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

func (r *PositionRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "UPDATE positions SET deleted_at = $1 WHERE position_id = $2 and deleted_at is null"

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
