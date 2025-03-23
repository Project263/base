package router

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"theaesthetics.ru/base/internal/handlers"
	"theaesthetics.ru/base/internal/repository"
	"theaesthetics.ru/base/internal/services"
)

func InitRouter(e *echo.Echo, db *pgxpool.Pool) {
	api := e.Group("/api/v1")

	eqRepo := repository.NewEquipmentRepository(db)
	eqService := services.NewEquipmentService(eqRepo)
	eqHandler := handlers.NewEquipmentHandler(eqService)

	api.GET("/equipment", eqHandler.GetAllEquipments)
	api.POST("/equipment", eqHandler.CreateEqipment)
	api.GET("/equipment/:id", eqHandler.GetEquipmentById)
	api.PUT("/equipment/:id", eqHandler.UpdateEqipment)
	api.DELETE("/equipment/:id", eqHandler.DeleteEquipment)

	msRepo := repository.NewMusclesRepository(db)
	msService := services.NewMusclesService(msRepo)
	msHandler := handlers.NewMusclesHandler(msService)

	api.GET("/muscles", msHandler.GetAllEMuscles)
	api.POST("/muscles", msHandler.CreateMuscles)
	api.GET("/muscles/:id", msHandler.GetMusclesById)
	api.PUT("/muscles/:id", msHandler.UpdateMuscles)
	api.DELETE("/muscles/:id", msHandler.DeleteMuscles)

	trRepo := repository.NewTrainsRepository(db)
	trService := services.NewTrainService(trRepo)
	trHandler := handlers.NewTrainHandler(trService)

	api.GET("/trains", trHandler.GetAllTrains)
	api.POST("/train", trHandler.CreateTrain)
}
