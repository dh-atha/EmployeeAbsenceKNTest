package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type AttendanceRepository interface {
	GetByID(ctx context.Context, id int) (*model.Attendance, error)
	Create(ctx context.Context, emp *model.Attendance) (int, error)
	Update(ctx context.Context, emp *model.Attendance) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.AttendanceList, error)
	GetReport(ctx context.Context, startDate, endDate string) ([]*model.AttendanceReport, error)
}
