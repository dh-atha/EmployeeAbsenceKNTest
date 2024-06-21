package service

import (
	"context"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
)

type LocationServiceInterface interface {
	Create(ctx context.Context, emp *model.Location) (int, error)
	Update(ctx context.Context, emp *model.Location) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.Location, error)
	GetAll(ctx context.Context) ([]*model.LocationList, error)
}

type LocationServiceImpl struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationServiceInterface {
	return &LocationServiceImpl{repo: repo}
}

func (s *LocationServiceImpl) Create(ctx context.Context, emp *model.Location) (int, error) {
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

func (s *LocationServiceImpl) Update(ctx context.Context, emp *model.Location) error {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()

	emp.UpdatedBy = &employeeName
	emp.UpdatedAt = &now

	return s.repo.Update(ctx, emp)
}

func (s *LocationServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *LocationServiceImpl) GetByID(ctx context.Context, id int) (*model.Location, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *LocationServiceImpl) GetAll(ctx context.Context) ([]*model.LocationList, error) {
	return s.repo.GetAll(ctx)
}
