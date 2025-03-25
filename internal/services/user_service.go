package services

import (
	"context"

	"theaesthetics.ru/base/internal/models"
	"theaesthetics.ru/base/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUserById(ctx context.Context, id uint) (models.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *UserService) GetUserAchievements(ctx context.Context, userId uint) ([]models.Achievenment, error) {
	return s.repo.GetUserAchievements(ctx, userId)
}
