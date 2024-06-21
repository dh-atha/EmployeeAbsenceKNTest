package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type PositionRepository interface {
	GetByID(ctx context.Context, id int) (*model.Position, error)
	Create(ctx context.Context, emp *model.Position) (int, error)
	Update(ctx context.Context, emp *model.Position) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.PositionList, error)
}
