package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.Employee
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.authService.GetEmployeeByCode(c, *req.EmployeeCode)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := h.authService.VerifyPassword(employee.Password, req.Password); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	data, err := h.authService.GenerateToken(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success login", data))
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.Employee
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Password) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password can't be empty"})
		return
	}

	employeeCode, err := h.authService.RegisterEmployee(c, &req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register employee"})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success register employee", map[string]interface{}{
		"employee_code": employeeCode,
	}))
}
