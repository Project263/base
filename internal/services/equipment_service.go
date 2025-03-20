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

func (s *EquipmentService) CreateEqipment(ctx context.Context, title, image string) error {
	return s.repo.CreateEqipment(ctx, title, image)
}

func (s *EquipmentService) RemoveEquipment(ctx context.Context, id uint8) error {
	return s.repo.RemoveEquipment(ctx, id)
}

func (s *EquipmentService) UpdateEquipment(ctx context.Context, equipment models.Equipment) error {
	return s.repo.UpdateEquipment(ctx, equipment)
}
