package handlers

import (
	"net/http"
	"strconv"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service service.DepartmentServiceInterface
}

func NewDepartmentHandler(DepartmentService service.DepartmentServiceInterface) *DepartmentHandler {
	return &DepartmentHandler{service: DepartmentService}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var emp model.Department
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DepartmentID, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create Department", gin.H{
		"department_id": DepartmentID,
	}))
}

func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Department ID"})
		return
	}

	var emp model.Department
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.DepartmentID = id

	if err := h.service.Update(c, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success update Department", nil))
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Department ID"})
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success delete department", nil))
}

func (h *DepartmentHandler) GetDepartmentByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	Department, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Department", Department))
}

func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	ctx := c.Request.Context()
	Departments, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Department", Departments))
}
