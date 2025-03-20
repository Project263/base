package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type EquipmentService struct {
	repo *repository.EquipmentRepository
}

func NewEquipmentService(repo *repository.EquipmentRepository) *EquipmentService {
	return &EquipmentService{repo: repo}
}

func (s *EquipmentService) GetAllEquipments(ctx context.Context) ([]models.Equipment, error) {
	return s.repo.GetAllEquipments(ctx)
}

func (s *EquipmentService) GetEquipmentById(ctx context.Context, id uint8) (*models.Equipment, error) {
	return s.repo.GetEquipmentById(ctx, id)
}
