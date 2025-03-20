package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"theaesthetics.ru/base/internal/services"
)

type EquipmentHandler struct {
	service *services.EquipmentService
}

func NewEquipmentHandler(service *services.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{service: service}
}

func (h *EquipmentHandler) GetAllEquipments(c echo.Context) error {
	equipments, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, equipments)
}
