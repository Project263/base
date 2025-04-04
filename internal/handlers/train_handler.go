package handlers

import (
	"base/internal/models"
	"base/internal/services"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TrainsHandler struct {
	service *services.TrainsService
}

func NewTrainsHandler(service *services.TrainsService) *TrainsHandler {
	return &TrainsHandler{service: service}
}

// Получить все тренировки
func (h *TrainsHandler) GetAllTrains(c echo.Context) error {
	ctx := c.Request().Context()
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		logrus.Error(err)
		return responseWithError(c, http.StatusBadRequest, errors.New("ошибка получения page"))
	}

	size, _ := strconv.Atoi(c.QueryParam("size"))

	if page < 1 {
		return responseWithError(c, http.StatusBadRequest, errors.New("page не может быть меньше 1"))
	}

	if size < 1 {
		size = 20
	}

	trains, total, err := h.service.GetAllTrains(ctx, page, size)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := map[string]interface{}{
		"data":       trains,
		"page":       page,
		"size":       size,
		"total":      total,
		"totalPages": (total + size - 1) / size,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TrainsHandler) GetTrainById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	train, err := h.service.GetTrainById(ctx, id)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, train)
}

func (h *TrainsHandler) CreateTrain(c echo.Context) error {
	ctx := context.Background()
	var train models.TrainWithMuscle
	err := c.Bind(&train)

	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.service.CreateTrain(ctx, train.Title, train.Description, train.Image, train.Video_url, train.Muscles.Id)

	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created",
	})
}

func (h *TrainsHandler) UpdateTrain(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")

	var train models.TrainWithMuscle
	train.Id = id
	err := c.Bind(&train)

	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.service.UpdateTrain(ctx, train)

	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "updated",
	})
}

func (h *TrainsHandler) DeleteTrain(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")

	err := h.service.DeleteTrain(ctx, id)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted",
	})
}
