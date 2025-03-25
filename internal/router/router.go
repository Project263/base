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

	ahRepo := repository.NewAchievementsRepository(db)
	ahService := services.NewAchievementsService(ahRepo)
	ahHandler := handlers.NewAhievementsHandler(ahService)

	api.GET("/ahievements", ahHandler.GetAllAchievements)
	api.POST("/ahievements", ahHandler.CreateAchievement)
	api.GET("/ahievements/:id", ahHandler.GetAchievementById)
	api.PUT("/ahievements/:id", ahHandler.UpdateAchievement)
	api.DELETE("/ahievements/:id", ahHandler.DeleteAchievement)

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
	api.GET("/trains/:id", trHandler.GetTrainById)

	exRepo := repository.NewExercisesRepository(db)
	exService := services.NewExercisesService(exRepo)
	exHandler := handlers.NewExerciseHandler(exService)

	api.GET("/exercises", exHandler.GetAllExercises)
	api.POST("/exercises", exHandler.CreateExercise)
	api.GET("/exercises/:id", exHandler.GetExerciseById)
	api.PUT("/exercises/:id", exHandler.UpdateExercise)
	api.DELETE("/exercises/:id", exHandler.DeleteExercise)

	usRepo := repository.NewUserRepository(db)
	usService := services.NewUserService(usRepo)
	usHandler := handlers.NewUserHandler(usService)

	api.GET("/users", usHandler.GetAllUsers)
	api.GET("/users/:id", usHandler.GetUserById)
}
