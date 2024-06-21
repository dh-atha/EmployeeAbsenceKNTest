package repository

import (
	"context"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
)

type LocationRepository interface {
	GetByID(ctx context.Context, id int) (*model.Location, error)
	Create(ctx context.Context, emp *model.Location) (int, error)
	Update(ctx context.Context, emp *model.Location) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.LocationList, error)
}
