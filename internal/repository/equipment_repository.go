package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"theaesthetics.ru/base/internal/models"
)

const tableNameEquipments = "equipments"

type EquipmentRepository struct {
	db *pgxpool.Pool
}

func NewEquipmentRepository(db *pgxpool.Pool) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) GetAllEquipments(ctx context.Context) ([]models.Equipment, error) {
	query := `SELECT id, title, image FROM equipments`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var equipments []models.Equipment
	for rows.Next() {
		var equipment models.Equipment
		err := rows.Scan(&equipment.Id, &equipment.Title, &equipment.Image)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, equipment)
	}
	return equipments, nil
}

func (r *EquipmentRepository) GetEquipmentById(ctx context.Context, id uint8) (*models.Equipment, error) {
	query := `SELECT id, title, image FROM equipments WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	equipment := &models.Equipment{}
	err := row.Scan(&equipment.Id, &equipment.Title, &equipment.Image)

	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (r *EquipmentRepository) CreateEqipment(ctx context.Context, title, image string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, title, tableNameEquipments)
	if err != nil {
		return err
	}

	query := `INSERT INTO equipments (title, image) VALUES ($1, $2)`
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

func (r *EquipmentRepository) DeleteEquipment(ctx context.Context, id uint8) error {
	query := `DELETE FROM equipments WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EquipmentRepository) UpdateEquipment(ctx context.Context, equipment models.Equipment) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		logrus.Error()
		return err
	}
	defer tx.Rollback(ctx)

	err = checkTitle(tx, ctx, equipment.Title, tableNameEquipments)
	if err != nil {
		return err
	}

	query := `UPDATE equipments SET title = $1, image = $2 WHERE id = $3`
	_, err = tx.Exec(ctx, query, equipment.Title, equipment.Image, equipment.Id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}
