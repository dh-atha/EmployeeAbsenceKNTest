package postgre

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type AttendanceRepositoryImpl struct {
	db *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) repository.AttendanceRepository {
	return &AttendanceRepositoryImpl{db: db}
}

func (r *AttendanceRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Attendance, error) {
	query := "SELECT * FROM attendances WHERE attendance_id = $1 and deleted_at is null"

	var emp model.Attendance
	err := r.db.GetContext(ctx, &emp, query, id)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *AttendanceRepositoryImpl) GetAll(ctx context.Context) ([]*model.AttendanceList, error) {
	var Attendances []*model.AttendanceList
	query := "SELECT attendance_id, employee_id, location_id, absent_in, absent_out FROM attendances WHERE deleted_at IS NULL"

	err := r.db.SelectContext(ctx, &Attendances, query)
	if err != nil {
		return nil, err
	}

	return Attendances, nil
}

func (r *AttendanceRepositoryImpl) Create(ctx context.Context, emp *model.Attendance) (int, error) {
	query := "INSERT INTO attendances (employee_id, location_id, absent_in, absent_out, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING attendance_id"

	var AttendanceID int
	err := r.db.QueryRowContext(ctx, query,
		emp.EmployeeID,
		emp.LocationID,
		emp.AbsentIn,
		emp.AbsentOut,
		emp.CreatedAt,
		emp.CreatedBy,
		emp.UpdatedAt,
		emp.UpdatedBy,
	).Scan(&AttendanceID)
	if err != nil {
		return 0, err
	}

	return AttendanceID, nil
}

func (r *AttendanceRepositoryImpl) Update(ctx context.Context, emp *model.Attendance) error {
	query := `
		UPDATE attendances SET
			employee_id = COALESCE(NULLIF(:employee_id, 0), employee_id),
			location_id = COALESCE(NULLIF(:location_id, 0), location_id),
			absent_in = COALESCE(:absent_in, absent_in),
			absent_out = COALESCE(:absent_out, absent_out),
			updated_at = :updated_at,
			updated_by = :updated_by
		WHERE attendance_id = :attendance_id AND deleted_at IS NULL
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

func (r *AttendanceRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := "UPDATE attendances SET deleted_at = $1 WHERE attendance_id = $2 and deleted_at is null"

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

func (r *AttendanceRepositoryImpl) GetReport(ctx context.Context, startDate, endDate string) ([]*model.AttendanceReport, error) {
	query := `
        SELECT
            a.absent_in,
            a.absent_out,
            e.employee_code,
            e.employee_name,
            d.department_name,
            p.position_name,
            l.location_name
        FROM
            attendances a
            INNER JOIN employees e ON a.employee_id = e.employee_id
            INNER JOIN departments d ON e.department_id = d.department_id
            INNER JOIN positions p ON e.position_id = p.position_id
            INNER JOIN locations l ON a.location_id = l.location_id
        WHERE
            a.absent_in BETWEEN $1::timestamp AND $2::timestamp
            AND a.deleted_at IS NULL
    `

	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []*model.AttendanceReport

	for rows.Next() {
		var r model.AttendanceReport
		err := rows.Scan(
			&r.AbsentIn,
			&r.AbsentOut,
			&r.EmployeeCode,
			&r.EmployeeName,
			&r.DepartmentName,
			&r.PositionName,
			&r.LocationName,
		)
		if err != nil {
			return nil, err
		}
		report = append(report, &r)
	}

	return report, nil
}
