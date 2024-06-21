package handlers

import (
	"net/http"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/security"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/interfaces/api/middleware"
	"github.com/dh-atha/EmployeeAbsenceKNTest/pkg/config"
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Data:    data,
	}
}

type Handlers struct {
	jwtService        security.JWTService
	authHandler       *AuthHandler
	employeeHandler   *EmployeeHandler
	departmentHandler *DepartmentHandler
	positionHandler   *PositionHandler
	locationHandler   *LocationHandler
	attendanceHandler *AttendanceHandler
}

type HandlersRequirements struct {
	JwtService        security.JWTService
	AuthHandler       *AuthHandler
	EmployeeHandler   *EmployeeHandler
	DepartmentHandler *DepartmentHandler
	PositionHandler   *PositionHandler
	LocationHandler   *LocationHandler
	AttendanceHandler *AttendanceHandler
}

func NewHandlers(req HandlersRequirements) *Handlers {
	return &Handlers{
		jwtService:        req.JwtService,
		authHandler:       req.AuthHandler,
		employeeHandler:   req.EmployeeHandler,
		departmentHandler: req.DepartmentHandler,
		positionHandler:   req.PositionHandler,
		locationHandler:   req.LocationHandler,
		attendanceHandler: req.AttendanceHandler,
	}
}

func (a *Handlers) CreateServer(address string) (*http.Server, error) {
	gin.SetMode(config.Configuration.Server.Mode)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/ping", a.checkConnectivity)

	r.POST("/login", a.authHandler.Login)
	r.POST("/register", a.authHandler.Register)
	r.Use(middleware.AuthMiddleware(a.jwtService))

	r.POST("/employee", a.employeeHandler.CreateEmployee)
	r.GET("/employee/:id", a.employeeHandler.GetEmployeeByID)
	r.GET("/employees", a.employeeHandler.GetAllEmployees)
	r.PUT("/employee/:id", a.employeeHandler.UpdateEmployee)
	r.DELETE("/employee/:id", a.employeeHandler.DeleteEmployee)

	r.POST("/department", a.departmentHandler.CreateDepartment)
	r.GET("/department/:id", a.departmentHandler.GetDepartmentByID)
	r.GET("/departments", a.departmentHandler.GetAllDepartments)
	r.PUT("/department/:id", a.departmentHandler.UpdateDepartment)
	r.DELETE("/department/:id", a.departmentHandler.DeleteDepartment)

	r.POST("/position", a.positionHandler.CreatePosition)
	r.GET("/position/:id", a.positionHandler.GetPositionByID)
	r.GET("/positions", a.positionHandler.GetAllPositions)
	r.PUT("/position/:id", a.positionHandler.UpdatePosition)
	r.DELETE("/position/:id", a.positionHandler.DeletePosition)

	r.POST("/location", a.locationHandler.CreateLocation)
	r.GET("/location/:id", a.locationHandler.GetLocationByID)
	r.GET("/locations", a.locationHandler.GetAllLocations)
	r.PUT("/location/:id", a.locationHandler.UpdateLocation)
	r.DELETE("/location/:id", a.locationHandler.DeleteLocation)

	r.POST("/attendance", a.attendanceHandler.CreateAttendance)
	r.GET("/attendance/:id", a.attendanceHandler.GetAttendanceByID)
	r.GET("/attendances", a.attendanceHandler.GetAllAttendances)
	r.PUT("/attendance/:id", a.attendanceHandler.UpdateAttendance)
	r.DELETE("/attendance/:id", a.attendanceHandler.DeleteAttendance)
	r.GET("/attendance/report", a.attendanceHandler.GetAttendanceReport)

	server := &http.Server{
		Addr:    address,
		Handler: r,
	}

	return server, nil
}

func (a *Handlers) checkConnectivity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
