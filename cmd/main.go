package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/domain/service"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/db/postgre"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/infrastructure/security"
	"github.com/dh-atha/EmployeeAbsenceKNTest/internal/interfaces/api/handlers"
	"github.com/dh-atha/EmployeeAbsenceKNTest/pkg/config"
	"github.com/jmoiron/sqlx"
)

func main() {
	config.LoadConfig("config.yaml")

	db, err := postgre.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	jwtService := security.NewJWTService(config.Configuration.Server.JWTSecret, "EmployeeAbsenceTest", config.Configuration.Server.TokenDuration)

	employeeRepo := postgre.NewEmployeeRepository(db)
	departmentRepo := postgre.NewDepartmentRepository(db)
	positionRepo := postgre.NewPositionRepository(db)
	locationRepo := postgre.NewLocationRepository(db)
	attendanceRepo := postgre.NewAttendanceRepository(db)

	authService := service.NewAuthService(employeeRepo, jwtService)
	employeeService := service.NewEmployeeService(employeeRepo)
	departmentService := service.NewDepartmentService(departmentRepo)
	positionService := service.NewPositionService(positionRepo)
	locationService := service.NewLocationService(locationRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo)

	authHander := handlers.NewAuthHandler(authService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	departmentHandler := handlers.NewDepartmentHandler(departmentService)
	positionHandler := handlers.NewPositionHandler(positionService)
	locationHandler := handlers.NewLocationHandler(locationService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)

	handlers := handlers.NewHandlers(handlers.HandlersRequirements{
		JwtService:        jwtService,
		AuthHandler:       authHander,
		EmployeeHandler:   employeeHandler,
		DepartmentHandler: departmentHandler,
		PositionHandler:   positionHandler,
		LocationHandler:   locationHandler,
		AttendanceHandler: attendanceHandler,
	})

	server, err := handlers.CreateServer(fmt.Sprintf(":%d", config.Configuration.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}

	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdownCh
	log.Printf("Receiving signal: %s", sig)

	ctx, cancel := context.WithTimeout(context.Background(), config.Configuration.Server.ShutdownTimeout)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	func(db *sqlx.DB) {
		db.Close()
		log.Fatal("Successfully disconnected from Postgres..")
		log.Fatal("Successfully shutdown server..")
	}(db)
}
