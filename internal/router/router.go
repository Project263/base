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
}
