package service

import (
	"context"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
)

type DepartmentServiceInterface interface {
	Create(ctx context.Context, emp *model.Department) (int, error)
	Update(ctx context.Context, emp *model.Department) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.Department, error)
	GetAll(ctx context.Context) ([]*model.DepartmentList, error)
}

type DepartmentServiceImpl struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentServiceInterface {
	return &DepartmentServiceImpl{repo: repo}
}

func (s *DepartmentServiceImpl) Create(ctx context.Context, emp *model.Department) (int, error) {
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

func (s *DepartmentServiceImpl) Update(ctx context.Context, emp *model.Department) error {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()

	emp.UpdatedBy = &employeeName
	emp.UpdatedAt = &now

	return s.repo.Update(ctx, emp)
}

func (s *DepartmentServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *DepartmentServiceImpl) GetByID(ctx context.Context, id int) (*model.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DepartmentServiceImpl) GetAll(ctx context.Context) ([]*model.DepartmentList, error) {
	return s.repo.GetAll(ctx)
}
