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

type MusclesHandler struct {
	service *services.MusclesService
}

func NewMusclesHandler(service *services.MusclesService) *MusclesHandler {
	return &MusclesHandler{service: service}
}

func (h *MusclesHandler) GetAllEMuscles(c echo.Context) error {
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

	muscles, total, err := h.service.GetAllMuscless(ctx, page, size)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := map[string]interface{}{
		"data":       muscles,
		"page":       page,
		"size":       size,
		"total":      total,
		"totalPages": (total + size - 1) / size,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MusclesHandler) GetMusclesById(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var idu8 uint8 = uint8(id)
	Muscles, err := h.service.GetMusclesById(ctx, idu8)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, Muscles)
}

func (h *MusclesHandler) CreateMuscles(c echo.Context) error {
	ctx := context.Background()
	var Muscles models.Muscle
	err := c.Bind(&Muscles)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = h.service.CreateEqipment(ctx, Muscles.Title, Muscles.Image)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created",
	})
}

func (h *MusclesHandler) DeleteMuscles(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	uid := uint8(id)
	err = h.service.DeleteMuscles(ctx, uid)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted",
	})
}

func (h *MusclesHandler) UpdateMuscles(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")

	Muscles := models.Muscle{Id: id}
	err := c.Bind(&Muscles)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = h.service.UpdateMuscles(ctx, Muscles)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "updated",
	})
}
