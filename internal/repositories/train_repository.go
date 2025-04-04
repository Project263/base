package repositories

import (
	"base/internal/models"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const tableNameTrains = "trains"

type TrainsRepository struct {
	db   *pgxpool.Pool
	psql sq.StatementBuilderType
}

func NewTrainsRepository(db *pgxpool.Pool) *TrainsRepository {
	return &TrainsRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *TrainsRepository) GetAllTrains(ctx context.Context, page, size int) ([]models.TrainWithMuscle, int, error) {
	offset := (page - 1) * size

	query, args, err := r.psql.
		Select("trains.id", "trains.title", "trains.description", "trains.image", "trains.video_url", "muscles.id", "muscles.title", "muscles.image").
		From(tableNameTrains).
		Join("muscles ON trains.muscle_id = muscles.id").
		Limit(uint64(size)).
		Offset(uint64(offset)).
		ToSql()

	logrus.Info(query)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var trains []models.TrainWithMuscle
	for rows.Next() {
		var train models.TrainWithMuscle
		var muscle models.Muscle
		if err := rows.Scan(&train.Id, &train.Title, &train.Description, &train.Image, &train.Video_url,
			&muscle.Id, &muscle.Title, &muscle.Image); err != nil {
			return nil, 0, err
		}
		train.Muscles = muscle
		trains = append(trains, train)
	}

	// Получаем общее количество записей
	query, args, err = r.psql.
		Select("COUNT(*)").
		From(tableNameTrains).
		Join("muscles ON trains.muscle_id = muscles.id").
		ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total int
	err = r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return trains, total, nil
}

func (r *TrainsRepository) GetTrainById(ctx context.Context, id string) (*models.TrainWithMuscle, error) {
	query, args, err := r.psql.
		Select("trains.id", "trains.title", "trains.description", "trains.image", "trains.video_url",
			"muscles.id AS muscle_id", "muscles.title AS muscle_title", "muscles.image AS muscle_image").
		From(tableNameTrains).
		Join("muscles AS muscles ON trains.muscle_id = muscles.id").
		Where(sq.Eq{"trains.id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, query, args...)

	train := &models.TrainWithMuscle{}
	var muscle models.Muscle
	err = row.Scan(&train.Id, &train.Title, &train.Description, &train.Image, &train.Video_url,
		&muscle.Id, &muscle.Title, &muscle.Image)
	if err != nil {
		return nil, err
	}
	train.Muscles = muscle

	return train, nil
}

func (r *TrainsRepository) CreateTrain(ctx context.Context, title, description, image, videoUrl string, muscleId int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query, args, err := r.psql.
		Insert(tableNameTrains).
		Columns("title", "description", "image", "video_url", "muscle_id").
		Values(title, description, image, videoUrl, muscleId).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

func (r *TrainsRepository) DeleteTrain(ctx context.Context, id string) error {
	query, args, err := r.psql.
		Delete(tableNameTrains).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *TrainsRepository) UpdateTrain(ctx context.Context, train models.TrainWithMuscle) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query, args, err := r.psql.
		Update(tableNameTrains).
		Set("title", train.Title).
		Set("description", train.Description).
		Set("image", train.Image).
		Set("video_url", train.Video_url).
		Set("muscle_id", train.Muscles.Id).
		Where(sq.Eq{"id": train.Id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
