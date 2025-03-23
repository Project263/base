package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// checkTitle проверяет уникальность названия в переданной таблице
func checkTitle(tx pgx.Tx, ctx context.Context, title string, table string) error {
	// Безопасно подставляем имя таблицы
	query := fmt.Sprintf(`SELECT id FROM %s WHERE title = $1 FOR UPDATE`, table)
	var id int
	logrus.Infof("Проверяем уникальность в таблице: %s", table)

	err := tx.QueryRow(ctx, query, title).Scan(&id)
	if err == nil {
		return fmt.Errorf("'%s' with title '%s' already exists", table, title)
	}

	if err != pgx.ErrNoRows {
		return err
	}

	return nil
}
