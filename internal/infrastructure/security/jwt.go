package security

import (
	"fmt"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(config.Configuration.Server.JWTSecret)

type JWTService interface {
	GenerateToken(employeeCode, employeeName string, employeeID int) (model.GenerateJWT, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey     string
	issuer        string
	tokenDuration time.Duration
}

func NewJWTService(secretKey, issuer string, tokenDuration time.Duration) JWTService {
	return &jwtService{
		secretKey:     secretKey,
		issuer:        issuer,
		tokenDuration: tokenDuration,
	}
}

func (j *jwtService) GenerateToken(employeeCode, employeeName string, employeeID int) (model.GenerateJWT, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	exp := time.Now().UTC().Add(j.tokenDuration)
	claims := token.Claims.(jwt.MapClaims)
	claims["employee_code"] = employeeCode
	claims["employee_name"] = employeeName
	claims["employee_id"] = employeeID
	claims["exp"] = exp.Unix()
	claims["iat"] = time.Now().UTC().Unix()
	claims["iss"] = j.issuer

	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return model.GenerateJWT{}, fmt.Errorf("error generating token: %w", err)
	}

	return model.GenerateJWT{
		Token:        tokenString,
		ExpiredAt:    exp.String(),
		EmployeeID:   employeeID,
		EmployeeName: employeeName,
		EmployeeCode: employeeCode,
	}, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	return token, nil
}
