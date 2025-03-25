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

type AchievementsHandler struct {
	service *services.AchievementsService
}

func NewAhievementsHandler(service *services.AchievementsService) *AchievementsHandler {
	return &AchievementsHandler{service: service}
}

func (h *AchievementsHandler) GetAllAchievements(c echo.Context) error {
	ctx := c.Request().Context()
	achievements, err := h.service.GetAllAchievements(ctx)
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, achievements)
}

func (h *AchievementsHandler) GetAchievementById(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var idu8 uint8 = uint8(id)
	achievement, err := h.service.GetAchievementById(ctx, idu8)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, achievement)
}

func (h *AchievementsHandler) CreateAchievement(c echo.Context) error {
	ctx := context.Background()
	var achieve models.Achievenment
	err := c.Bind(&achieve)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = h.service.CreateAchievement(ctx, achieve)

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

func (h *AchievementsHandler) DeleteAchievement(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	uid := uint8(id)
	err = h.service.DeleteAchievement(ctx, uid)

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

func (h *AchievementsHandler) UpdateAchievement(c echo.Context) error {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	uid := uint8(id)

	achieve := models.Achievenment{Id: uid}
	err = c.Bind(&achieve)

	if err != nil {
		logrus.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = h.service.UpdateAchievement(ctx, achieve)

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
