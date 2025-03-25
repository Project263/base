package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/services"
)

type ExerciseHandler struct {
	service *services.ExercisesService
}

func NewExerciseHandler(service *services.ExercisesService) *ExerciseHandler {
	return &ExerciseHandler{service: service}
}

func (h *ExerciseHandler) GetAllExercises(c echo.Context) error {
	ctx := context.Background()
	exercises, err := h.service.GetAllExercises(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to get exercises")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) GetExerciseById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid exercise ID")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint8(id)

	ctx := context.Background()
	exercise, err := h.service.GetExerciseById(ctx, uid)
	if err != nil {
		logrus.WithError(err).Error("Exercise not found")
		return respondWithError(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, exercise)
}

func (h *ExerciseHandler) CreateExercise(c echo.Context) error {
	ctx := context.Background()

	var exercise models.FullExercises
	if err := c.Bind(&exercise); err != nil {
		logrus.WithError(err).Error("Failed to bind request data")
		return respondWithError(c, http.StatusBadRequest, err)
	}

	id, err := h.service.CreateExercise(ctx, exercise)
	if err != nil {
		logrus.WithError(err).Error("Failed to create exercise")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Exercise created successfully",
		"id":      id,
	})
}

func (h *ExerciseHandler) UpdateExercise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid exercise ID")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint8(id)

	ctx := context.Background()

	var exercise models.FullExercises
	if err := c.Bind(&exercise); err != nil {
		logrus.WithError(err).Error("Failed to bind request data")
		return respondWithError(c, http.StatusBadRequest, err)
	}

	if err := h.service.UpdateExercise(ctx, uid, exercise); err != nil {
		logrus.WithError(err).Error("Failed to update exercise")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Exercise updated successfully",
	})
}

func (h *ExerciseHandler) DeleteExercise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Invalid exercise ID")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint8(id)

	ctx := context.Background()

	if err := h.service.DeleteExercise(ctx, uid); err != nil {
		logrus.WithError(err).Error("Failed to delete exercise")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Exercise deleted successfully",
	})
}
