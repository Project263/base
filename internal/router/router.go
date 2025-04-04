package router

import (
	"base/internal/handlers"
	"base/internal/repositories"
	"base/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, db *pgxpool.Pool) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := e.Group("/api")
	api.GET("/users", userHandler.GetAllUsers)

	musclesRepo := repositories.NewMusclesRepository(db)
	musclesService := services.NewMusclesService(musclesRepo)
	musclesHandler := handlers.NewMusclesHandler(musclesService)

	api.GET("/muscles", musclesHandler.GetAllEMuscles)
	api.POST("/muscles", musclesHandler.CreateMuscles)
	api.GET("/muscles/:id", musclesHandler.GetMusclesById)
	api.PUT("/muscles/:id", musclesHandler.UpdateMuscles)
	api.DELETE("/muscles/:id", musclesHandler.DeleteMuscles)

	trainsRepo := repositories.NewTrainsRepository(db)
	trainsService := services.NewTrainsService(trainsRepo)
	trainsHandler := handlers.NewTrainsHandler(trainsService)

	api.GET("/trains", trainsHandler.GetAllTrains)
	api.POST("/trains", trainsHandler.CreateTrain)
	api.GET("/trains/:id", trainsHandler.GetTrainById)
	api.PUT("/trains/:id", trainsHandler.UpdateTrain)
	api.DELETE("/trains/:id", trainsHandler.DeleteTrain)
}
