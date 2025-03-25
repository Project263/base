package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"theaesthetics.ru/base/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT 
			id, login, nickname, COALESCE(avatar, '') AS avatar, advanced_version, 
			COALESCE(phone, '') AS phone, is_verified_phone, email, is_verified_mail, 
			COALESCE(age, 0) AS age, COALESCE(height, 0) AS height, COALESCE(weight, 0) AS weight, 
			COALESCE(sex, '') AS sex, day_streak, is_train_today, points, created_at, updated_at
		FROM users
	`

	var users []models.User

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Login, &user.Nickname, &user.Avatar, &user.AdvancedVersion,
			&user.Phone, &user.IsVerifiedPhone, &user.Email, &user.IsVerifiedMail,
			&user.Age, &user.Height, &user.Weight, &user.Sex, &user.DayStreak,
			&user.IsTrainToday, &user.Points, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id uint) (models.User, error) {
	query := `
		SELECT 
			id, login, nickname, COALESCE(avatar, '') AS avatar, advanced_version, 
			COALESCE(phone, '') AS phone, is_verified_phone, email, is_verified_mail, 
			COALESCE(age, 0) AS age, COALESCE(height, 0) AS height, COALESCE(weight, 0) AS weight, 
			COALESCE(sex, '') AS sex, day_streak, is_train_today, points, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(
		&user.ID, &user.Login, &user.Nickname, &user.Avatar, &user.AdvancedVersion,
		&user.Phone, &user.IsVerifiedPhone, &user.Email, &user.IsVerifiedMail,
		&user.Age, &user.Height, &user.Weight, &user.Sex, &user.DayStreak,
		&user.IsTrainToday, &user.Points, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
