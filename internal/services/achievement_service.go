package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type AchievementsService struct {
	repo *repository.AchievementsRepository
}

func NewAchievementsService(repo *repository.AchievementsRepository) *AchievementsService {
	return &AchievementsService{repo: repo}
}

func (s *AchievementsService) GetAllAchievements(ctx context.Context) ([]models.Achievenment, error) {
	return s.repo.GetAllAchievements(ctx)
}

func (s *AchievementsService) GetAchievementById(ctx context.Context, id uint8) (*models.Achievenment, error) {
	return s.repo.GetAchievementById(ctx, id)
}

func (s *AchievementsService) CreateAchievement(ctx context.Context, achieve models.Achievenment) error {
	return s.repo.CreateAchievement(ctx, achieve)
}

func (s *AchievementsService) DeleteAchievement(ctx context.Context, id uint8) error {
	return s.repo.DeleteAchievement(ctx, id)
}

func (s *AchievementsService) UpdateAchievement(ctx context.Context, achieve models.Achievenment) error {
	return s.repo.UpdateAchievement(ctx, achieve)
}
