package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"theaesthetics.ru/base/internal/models"
)

const tableNameTrains = "trains"

type TrainsRepository struct {
	db *pgxpool.Pool
}

func NewTrainsRepository(db *pgxpool.Pool) *TrainsRepository {
	return &TrainsRepository{db: db}
}
func (r *TrainsRepository) GetAllTrains(ctx context.Context) ([]models.TrainWithMuscle, error) {
	query := `
		SELECT 
		trains.id, trains.title, description, trains.image,
		Video_url, difficult, duration_time, 
		muscles.id, muscles.title,muscles.image
		FROM trains
		JOIN muscles ON muscles.id = trains.lead_muscle_id
	`
	var trains []models.TrainWithMuscle

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var train models.TrainWithMuscle
		err := rows.Scan(&train.Id, &train.Title, &train.Description, &train.Image,
			&train.Video_url, &train.Difficult, &train.Duration_time, &train.Muscles.Id,
			&train.Muscles.Title, &train.Muscles.Image)
		if err != nil {
			return nil, err
		}

		trains = append(trains, train)
	}

	return trains, nil
}

func (r *TrainsRepository) GetTrainById(ctx context.Context, id uint8) (models.TrainWithMuscle, error) {
	query := `
		SELECT 
		trains.id, trains.title, description, trains.image,
		Video_url, difficult, duration_time, 
		muscles.id, muscles.title,muscles.image
		FROM trains
		JOIN muscles ON muscles.id = trains.lead_muscle_id
		WHERE trains.id = $1
	`
	var train models.TrainWithMuscle

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&train.Id, &train.Title, &train.Description, &train.Image,
		&train.Video_url, &train.Difficult, &train.Duration_time, &train.Muscles.Id,
		&train.Muscles.Title, &train.Muscles.Image)

	if err != nil {
		return models.TrainWithMuscle{}, err
	}

	return train, nil
}

func (r *TrainsRepository) CreateTrain(ctx context.Context, train models.Train) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err = checkTitle(tx, ctx, train.Title, tableNameTrains); err != nil {
		return err
	}

	if err = checkMuscles(tx, ctx, train.Lead_muscle_id); err != nil {
		return err
	}

	query := `INSERT INTO trains (title, description, image, video_url, difficult, duration_time, lead_muscle_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(ctx, query, train.Title, train.Description, train.Image, train.Video_url, train.Difficult, train.Duration_time, train.Lead_muscle_id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
