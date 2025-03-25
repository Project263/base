package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	ctx := context.Background()
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to get users")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint(id)

	ctx := context.Background()
	user, err := h.service.GetUserById(ctx, uid)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserAchievements(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid user ID")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint(id)

	ctx := context.Background()
	achievements, err := h.service.GetUserAchievements(ctx, uid)
	if err != nil {
		logrus.WithError(err).Error("Failed to get achievements")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, achievements)
}
