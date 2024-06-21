package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type DepartmentRepository interface {
	GetByID(ctx context.Context, id int) (*model.Department, error)
	Create(ctx context.Context, emp *model.Department) (int, error)
	Update(ctx context.Context, emp *model.Department) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.DepartmentList, error)
}
