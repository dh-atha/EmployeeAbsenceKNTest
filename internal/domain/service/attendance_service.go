package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
)

type AttendanceServiceInterface interface {
	Create(ctx context.Context, emp *model.Attendance) (int, error)
	Update(ctx context.Context, emp *model.Attendance) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.Attendance, error)
	GetAll(ctx context.Context) ([]*model.AttendanceList, error)
	GetReport(ctx context.Context, startDate, endDate time.Time) ([]*model.AttendanceReport, error)
}

type AttendanceServiceImpl struct {
	repo repository.AttendanceRepository
}

func NewAttendanceService(repo repository.AttendanceRepository) AttendanceServiceInterface {
	return &AttendanceServiceImpl{repo: repo}
}

func (s *AttendanceServiceImpl) Create(ctx context.Context, emp *model.Attendance) (int, error) {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()
	emp.CreatedAt = &now
	emp.CreatedBy = &employeeName
	emp.UpdatedAt = &now
	emp.UpdatedBy = &employeeName

	return s.repo.Create(ctx, emp)
}

func (s *AttendanceServiceImpl) Update(ctx context.Context, emp *model.Attendance) error {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()

	emp.UpdatedBy = &employeeName
	emp.UpdatedAt = &now

	return s.repo.Update(ctx, emp)
}

func (s *AttendanceServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *AttendanceServiceImpl) GetByID(ctx context.Context, id int) (*model.Attendance, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *AttendanceServiceImpl) GetAll(ctx context.Context) ([]*model.AttendanceList, error) {
	return s.repo.GetAll(ctx)
}

func (s *AttendanceServiceImpl) GetReport(ctx context.Context, startDate, endDate time.Time) ([]*model.AttendanceReport, error) {
	startDateStr := startDate.Format("2006-01-02 15:04:05")
	endDateStr := endDate.Format("2006-01-02 15:04:05")

	fmt.Println(startDateStr)
	fmt.Println(endDateStr)

	return s.repo.GetReport(ctx, startDateStr, endDateStr)
}
