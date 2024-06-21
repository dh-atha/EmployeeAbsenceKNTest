package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/constants"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/model"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	service service.AttendanceServiceInterface
}

func NewAttendanceHandler(AttendanceService service.AttendanceServiceInterface) *AttendanceHandler {
	return &AttendanceHandler{service: AttendanceService}
}

func (h *AttendanceHandler) CreateAttendance(c *gin.Context) {
	var emp model.Attendance
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if emp.EmployeeID == 0 {
		empId, _ := c.Get(constants.EmployeeIDJWTKey)
		empIdFloat := empId.(float64)
		emp.EmployeeID = int(empIdFloat)
	}

	if emp.AbsentIn == nil {
		now := time.Now().UTC()
		emp.AbsentIn = &now
	}

	AttendanceID, err := h.service.Create(c, &emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success create Attendance", gin.H{
		"attendance_id": AttendanceID,
	}))
}

func (h *AttendanceHandler) UpdateAttendance(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Attendance ID"})
		return
	}

	var emp model.Attendance
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.AttendanceID = id

	if err := h.service.Update(c, &emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success update Attendance", nil))
}

func (h *AttendanceHandler) DeleteAttendance(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Attendance ID"})
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success delete Attendance", nil))
}

func (h *AttendanceHandler) GetAttendanceByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Attendance ID"})
		return
	}

	Attendance, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Attendance", Attendance))
}

func (h *AttendanceHandler) GetAllAttendances(c *gin.Context) {
	ctx := c.Request.Context()
	Attendances, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get Attendance", Attendances))
}

func (h *AttendanceHandler) GetAttendanceReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02 15:04:05", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use yyyy-mm-dd hh:mm:ss"})
		return
	}

	endDate, err := time.Parse("2006-01-02 15:04:05", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use yyyy-mm-dd hh:mm:ss"})
		return
	}

	report, err := h.service.GetReport(c, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance report"})
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse("success get attendance report", report))
}
