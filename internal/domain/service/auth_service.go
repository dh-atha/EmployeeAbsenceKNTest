package service

import (
	"context"
	"errors"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/repository"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/encrypter"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/security"
)

type AuthServiceInterface interface {
	GetEmployeeByCode(ctx context.Context, code string) (*model.Employee, error)
	RegisterEmployee(ctx context.Context, employee *model.Employee) (string, error)
	VerifyPassword(hashedPassword, password string) error
	GenerateToken(employee *model.Employee) (model.GenerateJWT, error)
}

type AuthService struct {
	employeeRepo repository.EmployeeRepository
	jwtService   security.JWTService
}

func NewAuthService(employeeRepo repository.EmployeeRepository, jwtService security.JWTService) AuthServiceInterface {
	return &AuthService{
		employeeRepo: employeeRepo,
		jwtService:   jwtService,
	}
}

func (s *AuthService) GetEmployeeByCode(ctx context.Context, code string) (*model.Employee, error) {
	return s.employeeRepo.GetByCode(ctx, code)
}

func (s *AuthService) RegisterEmployee(ctx context.Context, employee *model.Employee) (string, error) {
	hashedPassword, err := encrypter.EncryptPassword(employee.Password, 10)
	if err != nil {
		return "", errors.Join(constants.ErrService, err)
	}
	employee.Password = hashedPassword

	now := time.Now().UTC()
	createdBy := "system"
	employee.CreatedAt = &now
	employee.UpdatedAt = &now
	employee.UpdatedBy = &createdBy
	employee.CreatedBy = &createdBy

	return s.employeeRepo.Create(ctx, employee)
}

func (s *AuthService) VerifyPassword(hashedPassword, password string) error {
	return encrypter.VerifyPassword(password, hashedPassword)
}

func (s *AuthService) GenerateToken(employee *model.Employee) (model.GenerateJWT, error) {
	return s.jwtService.GenerateToken(*employee.EmployeeCode, *employee.EmployeeName, employee.EmployeeID)
}
