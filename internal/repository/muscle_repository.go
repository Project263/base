package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"theaesthetics.ru/base/internal/models"
)

const tableNameMuscles = "muscles"

type MusclesRepository struct {
	db *pgxpool.Pool
}

func NewMusclesRepository(db *pgxpool.Pool) *MusclesRepository {
	return &MusclesRepository{db: db}
}

func (r *MusclesRepository) GetAllMuscles(ctx context.Context) ([]models.Muscles, error) {
	query := `SELECT id, title, image FROM muscles`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Muscless []models.Muscles
	for rows.Next() {
		var Muscles models.Muscles
		err := rows.Scan(&Muscles.Id, &Muscles.Title, &Muscles.Image)
		if err != nil {
			return nil, err
		}
		Muscless = append(Muscless, Muscles)
	}
	return Muscless, nil
}

func (r *MusclesRepository) GetMusclesById(ctx context.Context, id uint8) (*models.Muscles, error) {
	query := `SELECT id, title, image FROM muscles WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	Muscles := &models.Muscles{}
	err := row.Scan(&Muscles.Id, &Muscles.Title, &Muscles.Image)

	if err != nil {
		return nil, err
	}

	return Muscles, nil
}

func (r *MusclesRepository) CreateMuscles(ctx context.Context, title, image string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, title, tableNameMuscles)
	if err != nil {
		return err
	}

	query := `INSERT INTO muscles (title, image) VALUES ($1, $2)`
	_, err = r.db.Exec(ctx, query, title, image)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

func (r *MusclesRepository) DeleteMuscles(ctx context.Context, id uint8) error {
	query := `DELETE FROM muscles WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *MusclesRepository) UpdateMuscles(ctx context.Context, Muscles models.Muscles) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, Muscles.Title, tableNameMuscles)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `UPDATE muscles SET title = $1, image = $2 WHERE id = $3`, Muscles.Title, Muscles.Image, Muscles.Id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
