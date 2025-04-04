package repositories

import (
	"base/internal/models"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const tableNameMuscles = "muscles"

type MusclesRepository struct {
	db   *pgxpool.Pool
	psql sq.StatementBuilderType
}

func NewMusclesRepository(db *pgxpool.Pool) *MusclesRepository {
	return &MusclesRepository{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *MusclesRepository) GetAllMuscles(ctx context.Context, page, size int) ([]models.Muscle, int, error) {
	offset := (page - 1) * size

	query, args, err := r.psql.
		Select("id", "title", "image").
		From(tableNameMuscles).
		Limit(uint64(size)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var muscles []models.Muscle
	for rows.Next() {
		var muscle models.Muscle
		if err := rows.Scan(&muscle.Id, &muscle.Title, &muscle.Image); err != nil {
			return nil, 0, err
		}
		muscles = append(muscles, muscle)
	}

	query, args, err = r.psql.
		Select("COUNT(*)").
		From(tableNameMuscles).
		ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total int
	err = r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return muscles, total, nil
}

func (r *MusclesRepository) GetMusclesById(ctx context.Context, id uint8) (*models.Muscle, error) {
	query, args, err := r.psql.
		Select("id", "title", "image").
		From(tableNameMuscles).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, query, args...)

	muscle := &models.Muscle{}
	err = row.Scan(&muscle.Id, &muscle.Title, &muscle.Image)
	if err != nil {
		return nil, err
	}

	return muscle, nil
}

func (r *MusclesRepository) CreateMuscles(ctx context.Context, title, image string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err != nil {
		return err
	}

	query, args, err := r.psql.
		Insert(tableNameMuscles).
		Columns("title", "image").
		Values(title, image).
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

func (r *MusclesRepository) DeleteMuscles(ctx context.Context, id uint8) error {
	query, args, err := r.psql.
		Delete(tableNameMuscles).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	return err
}

func (r *MusclesRepository) UpdateMuscles(ctx context.Context, muscle models.Muscle) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err != nil {
		return err
	}

	query, args, err := r.psql.
		Update(tableNameMuscles).
		Set("title", muscle.Title).
		Set("image", muscle.Image).
		Where(sq.Eq{"id": muscle.Id}).
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
