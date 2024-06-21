package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type EmployeeRepository interface {
	GetByID(ctx context.Context, id int) (*model.Employee, error)
	GetByCode(ctx context.Context, id string) (*model.Employee, error)
	Create(ctx context.Context, emp *model.Employee) (string, error)
	Update(ctx context.Context, emp *model.Employee) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.EmployeeList, error)
}
