package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type DepartmentRepositoryImpl struct {
	db *sqlx.DB
}

func NewDepartmentRepository(db *sqlx.DB) repository.DepartmentRepository {
	return &DepartmentRepositoryImpl{db: db}
}

func (r *DepartmentRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Department, error) {
	query := "SELECT * FROM departments WHERE department_id = $1 and deleted_at is null"

	var emp model.Department
	err := r.db.GetContext(ctx, &emp, query, id)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *DepartmentRepositoryImpl) GetAll(ctx context.Context) ([]*model.DepartmentList, error) {
	var departments []*model.DepartmentList
	query := "SELECT department_id, department_name FROM departments WHERE deleted_at IS NULL"

	err := r.db.SelectContext(ctx, &departments, query)
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepositoryImpl) Create(ctx context.Context, emp *model.Department) (int, error) {
	query := "INSERT INTO departments (department_name, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5) RETURNING department_id"

	var departmentID int
	err := r.db.QueryRowContext(ctx, query,
		emp.DepartmentName,
		emp.CreatedAt,
		emp.CreatedBy,
		emp.UpdatedAt,
		emp.UpdatedBy,
	).Scan(&departmentID)
	if err != nil {
		return 0, err
	}

	return departmentID, nil
}

func (r *DepartmentRepositoryImpl) Update(ctx context.Context, emp *model.Department) error {
	query := `
		UPDATE departments SET
			Department_name = COALESCE(NULLIF(:department_name, ''), department_name),
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE department_id = :department_id AND deleted_at IS NULL
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

func (r *DepartmentRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "UPDATE departments SET deleted_at = $1 WHERE department_id = $2 and deleted_at is null"

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
