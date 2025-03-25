package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"theaesthetics.ru/base/internal/models"
)

type ExercisesRepository struct {
	db *pgxpool.Pool
}

func NewExercisesRepository(db *pgxpool.Pool) *ExercisesRepository {
	return &ExercisesRepository{db: db}
}
func (r *ExercisesRepository) GetAllExercises(ctx context.Context) ([]models.FullExercises, error) {
	query := `
		SELECT 
		exercises.id, 
		exercises.title, 
		COALESCE(exercises.description, '') AS description, 
		COALESCE(exercises.image, '') AS image, 
		COALESCE(exercises.video_url, '') AS video_url, 
		exercises.sets, 
		exercises.reps, 
		exercises.difficult, 
		COALESCE(equipments.id, 0) AS equipment_id, 
		COALESCE(equipments.title, '') AS equipment_title, 
		COALESCE(equipments.image, '') AS equipment_image, 
		COALESCE(muscles.id, 0) AS muscle_id, 
		COALESCE(muscles.title, '') AS muscle_title, 
		COALESCE(muscles.image, '') AS muscle_image
	FROM exercises
	LEFT JOIN equipments ON equipments.id = exercises.equipment_id
	LEFT JOIN muscles ON muscles.id = exercises.lead_muscle_id
	`
	var exercises []models.FullExercises

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var exercise models.FullExercises
		err := rows.Scan(&exercise.Id, &exercise.Title, &exercise.Description, &exercise.Image, &exercise.Video_url,
			&exercise.Sets, &exercise.Reps, &exercise.Difficult,
			&exercise.Equipment.Id, &exercise.Equipment.Title, &exercise.Equipment.Image,
			&exercise.Muscles.Id, &exercise.Muscles.Title, &exercise.Muscles.Image)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

func (r *ExercisesRepository) GetExerciseById(ctx context.Context, id uint8) (models.FullExercises, error) {
	query := `
		SELECT 
		exercises.id, 
		exercises.title, 
		COALESCE(exercises.description, '') AS description, 
		COALESCE(exercises.image, '') AS image, 
		COALESCE(exercises.video_url, '') AS video_url, 
		exercises.sets, 
		exercises.reps, 
		exercises.difficult, 
		COALESCE(equipments.id, 0) AS equipment_id, 
		COALESCE(equipments.title, '') AS equipment_title, 
		COALESCE(equipments.image, '') AS equipment_image, 
		COALESCE(muscles.id, 0) AS muscle_id, 
		COALESCE(muscles.title, '') AS muscle_title, 
		COALESCE(muscles.image, '') AS muscle_image
	FROM exercises
	LEFT JOIN equipments ON equipments.id = exercises.equipment_id
	LEFT JOIN muscles ON muscles.id = exercises.lead_muscle_id
	WHERE exercises.id = $1;

	`
	var exercise models.FullExercises

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&exercise.Id, &exercise.Title, &exercise.Description, &exercise.Image, &exercise.Video_url,
		&exercise.Sets, &exercise.Reps, &exercise.Difficult,
		&exercise.Equipment.Id, &exercise.Equipment.Title, &exercise.Equipment.Image,
		&exercise.Muscles.Id, &exercise.Muscles.Title, &exercise.Muscles.Image)

	if err != nil {
		return models.FullExercises{}, err
	}

	return exercise, nil
}

func (r *ExercisesRepository) CreateExercise(ctx context.Context, exercise models.FullExercises) (uint8, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO exercises (title, description, image, video_url, sets, reps, difficult, equipment_id, lead_muscle_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 
			(SELECT id FROM equipments WHERE title = $8), 
			(SELECT id FROM muscles WHERE title = $9))
		RETURNING id
	`

	var id uint64
	err = tx.QueryRow(ctx, query, exercise.Title, exercise.Description, exercise.Image, exercise.Video_url,
		exercise.Sets, exercise.Reps, exercise.Difficult, exercise.Equipment.Title, exercise.Muscles.Title).Scan(&id)

	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("transaction commit failed: %w", err)
	}

	return uint8(id), nil
}

func (r *ExercisesRepository) DeleteExercise(ctx context.Context, id uint8) error {
	query := `DELETE FROM exercises WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExercisesRepository) UpdateExercise(ctx context.Context, id uint8, exercise models.FullExercises) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE exercises 
		SET title = $1, description = $2, image = $3, video_url = $4, 
			sets = $5, reps = $6, difficult = $7, 
			equipment_id = (SELECT id FROM equipments WHERE title = $8), 
			lead_muscle_id = (SELECT id FROM muscles WHERE title = $9)
		WHERE id = $10
	`

	_, err = tx.Exec(ctx, query, exercise.Title, exercise.Description, exercise.Image, exercise.Video_url,
		exercise.Sets, exercise.Reps, exercise.Difficult, exercise.Equipment.Title, exercise.Muscles.Title, id)

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}
	return err
}
