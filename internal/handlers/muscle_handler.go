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

type MusclesHandler struct {
	service *services.MusclesService
}

func NewMusclesHandler(service *services.MusclesService) *MusclesHandler {
	return &MusclesHandler{service: service}
}

func (h *MusclesHandler) GetAllEMuscles(c echo.Context) error {
	ctx := c.Request().Context()
	Muscless, err := h.service.GetAllMuscless(ctx)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Muscless)
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
	var Muscles models.Muscles
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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	uid := uint8(id)

	Muscles := models.Muscles{Id: uid}
	err = c.Bind(&Muscles)

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
