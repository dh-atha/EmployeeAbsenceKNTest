package service

import (
	"context"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
)

type PositionServiceInterface interface {
	Create(ctx context.Context, emp *model.Position) (int, error)
	Update(ctx context.Context, emp *model.Position) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.Position, error)
	GetAll(ctx context.Context) ([]*model.PositionList, error)
}

type PositionServiceImpl struct {
	repo repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionServiceInterface {
	return &PositionServiceImpl{repo: repo}
}

func (s *PositionServiceImpl) Create(ctx context.Context, emp *model.Position) (int, error) {
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

func (s *PositionServiceImpl) Update(ctx context.Context, emp *model.Position) error {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()

	emp.UpdatedBy = &employeeName
	emp.UpdatedAt = &now

	return s.repo.Update(ctx, emp)
}

func (s *PositionServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *PositionServiceImpl) GetByID(ctx context.Context, id int) (*model.Position, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *PositionServiceImpl) GetAll(ctx context.Context) ([]*model.PositionList, error) {
	return s.repo.GetAll(ctx)
}
