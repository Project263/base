package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/internal/services"
)

type EquipmentHandler struct {
	service *services.EquipmentService
}

func NewEquipmentHandler(service *services.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{service: service}
}

func (h *EquipmentHandler) GetAllEquipments(c echo.Context) error {
	equipments, err := h.service.GetAllEquipments(c.Request().Context())
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, equipments)
}

func (h *EquipmentHandler) GetEquipmentById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var idu8 uint8 = uint8(id)
	equipment, err := h.service.GetEquipmentById(ctx, idu8)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, equipment)
}
