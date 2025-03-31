package handlers

import (
	"base/internal/services"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	ctx := context.Background()
	users, err := h.service.GetUsers(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to get users")
		return responseWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}
