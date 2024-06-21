package handlers

import (
	"net/http"
	"strconv"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service service.LocationServiceInterface
}

func NewLocationHandler(LocationService service.LocationServiceInterface) *LocationHandler {
	return &LocationHandler{service: LocationService}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var emp model.Location
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	LocationID, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create Location", gin.H{
		"location_id": LocationID,
	}))
}

func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Location ID"})
		return
	}

	var emp model.Location
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.LocationID = id

	if err := h.service.Update(c, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success update Location", nil))
}

func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Location ID"})
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success delete Location", nil))
}

func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Location ID"})
		return
	}

	Location, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Location", Location))
}

func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	ctx := c.Request.Context()
	Locations, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Location", Locations))
}
