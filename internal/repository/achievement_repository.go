package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/internal/models"
)

const tableNameAchievement = "achievements"

type AchievementsRepository struct {
	db *pgxpool.Pool
}

func NewAchievementsRepository(db *pgxpool.Pool) *AchievementsRepository {
	return &AchievementsRepository{db: db}
}

func (r *AchievementsRepository) GetAllAchievements(ctx context.Context) ([]models.Achievenment, error) {
	query := `SELECT id, title, image, description FROM achievements`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var achievements []models.Achievenment
	for rows.Next() {
		var achieve models.Achievenment
		err := rows.Scan(&achieve.Id, &achieve.Title, &achieve.Image, &achieve.Description)
		if err != nil {
			return nil, err
		}
		achievements = append(achievements, achieve)
	}
	return achievements, nil
}

func (r *AchievementsRepository) GetAchievementById(ctx context.Context, id uint8) (*models.Achievenment, error) {
	query := `SELECT id, title, image, description FROM achievements WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	achieve := &models.Achievenment{}
	err := row.Scan(&achieve.Id, &achieve.Title, &achieve.Image, &achieve.Description)

	if err != nil {
		return nil, err
	}

	return achieve, nil
}

func (r *AchievementsRepository) CreateAchievement(ctx context.Context, achieve models.Achievenment) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, achieve.Title, tableNameAchievement)
	if err != nil {
		return err
	}

	query := `INSERT INTO achievements (title, image, description) VALUES ($1, $2, $3)`
	_, err = r.db.Exec(ctx, query, achieve.Title, achieve.Image, achieve.Description)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

func (r *AchievementsRepository) DeleteAchievement(ctx context.Context, id uint8) error {
	query := `DELETE FROM achievements WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AchievementsRepository) UpdateAchievement(ctx context.Context, achieve models.Achievenment) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		logrus.Error()
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, achieve.Title, tableNameAchievement)
	if err != nil {
		return err
	}

	query := `UPDATE achievements SET title = $1, image = $2, description = $3 WHERE id = $4`
	_, err = r.db.Exec(ctx, query, achieve.Title, achieve.Image, achieve.Description, achieve.Id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
