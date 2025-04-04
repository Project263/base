package services

import (
	"base/internal/models"
	"base/internal/repositories"
	"context"
)

type MusclesService struct {
	repo *repositories.MusclesRepository
}

func NewMusclesService(repo *repositories.MusclesRepository) *MusclesService {
	return &MusclesService{repo: repo}
}

func (s *MusclesService) GetAllMuscless(ctx context.Context, page, size int) ([]models.Muscle, int, error) {
	return s.repo.GetAllMuscles(ctx, page, size)
}

func (s *MusclesService) GetMusclesById(ctx context.Context, id uint8) (*models.Muscle, error) {
	return s.repo.GetMusclesById(ctx, id)
}

func (s *MusclesService) CreateEqipment(ctx context.Context, title, image string) error {
	return s.repo.CreateMuscles(ctx, title, image)
}

func (s *MusclesService) DeleteMuscles(ctx context.Context, id uint8) error {
	return s.repo.DeleteMuscles(ctx, id)
}

func (s *MusclesService) UpdateMuscles(ctx context.Context, Muscles models.Muscle) error {
	return s.repo.UpdateMuscles(ctx, Muscles)
}
