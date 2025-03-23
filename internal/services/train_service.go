package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type TrainService struct {
	repo *repository.TrainsRepository
}

func NewTrainService(repo *repository.TrainsRepository) *TrainService {
	return &TrainService{repo: repo}
}

func (s *TrainService) CreateTrain(ctx context.Context, train models.Train) error {
	return s.repo.CreateTrain(ctx, train)
}

func (s *TrainService) GetAllTrains(ctx context.Context) ([]models.TrainWithMuscle, error) {
	return s.repo.GetAllTrains(ctx)
}

func (s *TrainService) GetTrainById(ctx context.Context, id uint8) (models.TrainWithMuscle, error) {
	return s.repo.GetTrainById(ctx, id)
}
