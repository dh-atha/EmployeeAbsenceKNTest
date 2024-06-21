package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepositoryImpl struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) repository.EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (r *EmployeeRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	query := "SELECT * FROM employees WHERE employee_id = $1 and deleted_at is null"

	var emp model.Employee
	err := r.db.GetContext(ctx, &emp, query, id)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}
func (r *EmployeeRepositoryImpl) GetByCode(ctx context.Context, code string) (*model.Employee, error) {
	query := "SELECT * FROM employees WHERE employee_code = $1 and deleted_at is null"

	var emp model.Employee
	err := r.db.GetContext(ctx, &emp, query, code)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepositoryImpl) GetAll(ctx context.Context) ([]*model.EmployeeList, error) {
	var employees []*model.EmployeeList
	query := "SELECT employee_id, employee_code, employee_name FROM employees WHERE deleted_at IS NULL"

	err := r.db.SelectContext(ctx, &employees, query)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepositoryImpl) Create(ctx context.Context, emp *model.Employee) (string, error) {
	query := "INSERT INTO employees (employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING employee_code"

	var employeeCode string
	err := r.db.QueryRowContext(ctx, query,
		emp.EmployeeName,
		emp.Password,
		emp.DepartmentID,
		emp.PositionID,
		emp.Superior,
		emp.CreatedAt,
		emp.CreatedBy,
		emp.UpdatedAt,
		emp.UpdatedBy,
	).Scan(&employeeCode)
	if err != nil {
		return "", err
	}

	return employeeCode, nil
}

func (r *EmployeeRepositoryImpl) Update(ctx context.Context, emp *model.Employee) error {
	query := `
		UPDATE employees SET
			employee_code = COALESCE(NULLIF(:employee_code, ''), employee_code),
			employee_name = COALESCE(NULLIF(:employee_name, ''), employee_name),
			department_id = COALESCE(NULLIF(:department_id, 0), department_id),
			position_id = COALESCE(NULLIF(:position_id, 0), position_id),
			superior = COALESCE(NULLIF(:superior, 0), superior),
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE employee_id = :employee_id AND deleted_at IS NULL
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

func (r *EmployeeRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "UPDATE employees SET deleted_at = $1 WHERE employee_id = $2 and deleted_at is null"

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
