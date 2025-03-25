package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type ExercisesService struct {
	repo *repository.ExercisesRepository
}

func NewExercisesService(repo *repository.ExercisesRepository) *ExercisesService {
	return &ExercisesService{repo: repo}
}

func (s *ExercisesService) GetAllExercises(ctx context.Context) ([]models.FullExercises, error) {
	return s.repo.GetAllExercises(ctx)
}

func (s *ExercisesService) GetExerciseById(ctx context.Context, id uint8) (models.FullExercises, error) {
	return s.repo.GetExerciseById(ctx, id)
}

func (s *ExercisesService) CreateExercise(ctx context.Context, exercise models.FullExercises) (uint8, error) {
	return s.repo.CreateExercise(ctx, exercise)
}

func (s *ExercisesService) UpdateExercise(ctx context.Context, id uint8, exercise models.FullExercises) error {
	return s.repo.UpdateExercise(ctx, id, exercise)
}

func (s *ExercisesService) DeleteExercise(ctx context.Context, id uint8) error {
	return s.repo.DeleteExercise(ctx, id)
}
