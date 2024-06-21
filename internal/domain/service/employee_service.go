package service

import (
	"context"
	"errors"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/encrypter"
)

type EmployeeServiceInterface interface {
	Create(ctx context.Context, emp *model.Employee) (string, error)
	Update(ctx context.Context, emp *model.Employee) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.Employee, error)
	GetAll(ctx context.Context) ([]*model.EmployeeList, error)
}

type EmployeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeServiceInterface {
	return &EmployeeServiceImpl{repo: repo}
}

func (s *EmployeeServiceImpl) Create(ctx context.Context, emp *model.Employee) (string, error) {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	hashed, err := encrypter.EncryptPassword(emp.Password, 10)
	if err != nil {
		return "", errors.Join(constants.ErrService, err)
	}

	emp.Password = hashed
	now := time.Now().UTC()
	emp.CreatedAt = &now
	emp.CreatedBy = &employeeName
	emp.UpdatedAt = &now
	emp.UpdatedBy = &employeeName

	return s.repo.Create(ctx, emp)
}

func (s *EmployeeServiceImpl) Update(ctx context.Context, emp *model.Employee) error {
	var (
		employeeName, _ = ctx.Value(constants.EmployeeNameJWTKey).(string)
	)
	now := time.Now().UTC()

	emp.UpdatedBy = &employeeName
	emp.UpdatedAt = &now

	return s.repo.Update(ctx, emp)
}

func (s *EmployeeServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *EmployeeServiceImpl) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeServiceImpl) GetAll(ctx context.Context) ([]*model.EmployeeList, error) {
	return s.repo.GetAll(ctx)
}
