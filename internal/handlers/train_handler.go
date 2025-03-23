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

type TrainHandler struct {
	service *services.TrainService
}

func NewTrainHandler(service *services.TrainService) *TrainHandler {
	return &TrainHandler{service: service}
}

func (h *TrainHandler) GetAllTrains(c echo.Context) error {
	ctx := context.Background()
	trains, err := h.service.GetAllTrains(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to get trains")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, trains)
}

func (h *TrainHandler) GetTrainById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to get trains")
		return respondWithError(c, http.StatusBadRequest, err)
	}
	uid := uint8(id)

	ctx := context.Background()
	trains, err := h.service.GetTrainById(ctx, uid)
	if err != nil {
		logrus.WithError(err).Error("Failed to get trains")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, trains)
}

func (h *TrainHandler) CreateTrain(c echo.Context) error {
	ctx := context.Background()

	var train models.Train
	if err := c.Bind(&train); err != nil {
		logrus.WithError(err).Error("Failed to bind request data")
		return respondWithError(c, http.StatusBadRequest, err)
	}

	if err := h.service.CreateTrain(ctx, train); err != nil {
		logrus.WithError(err).Error("Failed to create train")
		return respondWithError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Train created successfully",
	})
}

func respondWithError(c echo.Context, statusCode int, err error) error {
	return c.JSON(statusCode, map[string]string{
		"error": err.Error(),
	})
}
