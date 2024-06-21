package handlers

import (
	"net/http"
	"strconv"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	service service.PositionServiceInterface
}

func NewPositionHandler(PositionService service.PositionServiceInterface) *PositionHandler {
	return &PositionHandler{service: PositionService}
}

func (h *PositionHandler) CreatePosition(c *gin.Context) {
	var emp model.Position
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	PositionID, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create Position", gin.H{
		"position_id": PositionID,
	}))
}

func (h *PositionHandler) UpdatePosition(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Position ID"})
		return
	}

	var emp model.Position
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.PositionID = id

	if err := h.service.Update(c, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success update Position", nil))
}

func (h *PositionHandler) DeletePosition(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Position ID"})
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success delete Position", nil))
}

func (h *PositionHandler) GetPositionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Position ID"})
		return
	}

	Position, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Position", Position))
}

func (h *PositionHandler) GetAllPositions(c *gin.Context) {
	ctx := c.Request.Context()
	Positions, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Position", Positions))
}
