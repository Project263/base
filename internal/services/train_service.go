package services

import (
	"base/internal/models"
	"base/internal/repositories"
	"context"
)

type TrainsService struct {
	repo *repositories.TrainsRepository
}

func NewTrainsService(repo *repositories.TrainsRepository) *TrainsService {
	return &TrainsService{repo: repo}
}

func (s *TrainsService) GetAllTrains(ctx context.Context, page, size int) ([]models.TrainWithMuscle, int, error) {
	return s.repo.GetAllTrains(ctx, page, size)
}

func (s *TrainsService) GetTrainById(ctx context.Context, id string) (*models.TrainWithMuscle, error) {
	return s.repo.GetTrainById(ctx, id)
}

func (s *TrainsService) CreateTrain(ctx context.Context, title, description, image, videoUrl, muscleId string) error {
	return s.repo.CreateTrain(ctx, title, description, image, videoUrl, muscleId)
}

func (s *TrainsService) DeleteTrain(ctx context.Context, id string) error {
	return s.repo.DeleteTrain(ctx, id)
}

func (s *TrainsService) UpdateTrain(ctx context.Context, train models.TrainWithMuscle) error {
	return s.repo.UpdateTrain(ctx, train)
}
