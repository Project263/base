package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"theaesthetics.ru/base/internal/models"
)

type EquipmentRepository struct {
	db *pgxpool.Pool
}

func NewEquipmentRepository(db *pgxpool.Pool) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) GetAll(ctx context.Context) ([]models.Equipment, error) {
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
