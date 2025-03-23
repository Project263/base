package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type MusclesService struct {
	repo *repository.MusclesRepository
}

func NewMusclesService(repo *repository.MusclesRepository) *MusclesService {
	return &MusclesService{repo: repo}
}

func (s *MusclesService) GetAllMuscless(ctx context.Context) ([]models.Muscles, error) {
	return s.repo.GetAllMuscles(ctx)
}

func (s *MusclesService) GetMusclesById(ctx context.Context, id uint8) (*models.Muscles, error) {
	return s.repo.GetMusclesById(ctx, id)
}

func (s *MusclesService) CreateEqipment(ctx context.Context, title, image string) error {
	return s.repo.CreateMuscles(ctx, title, image)
}

func (s *MusclesService) DeleteMuscles(ctx context.Context, id uint8) error {
	return s.repo.DeleteMuscles(ctx, id)
}

func (s *MusclesService) UpdateMuscles(ctx context.Context, Muscles models.Muscles) error {
	return s.repo.UpdateMuscles(ctx, Muscles)
}
