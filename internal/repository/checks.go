package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func checkTitle(tx pgx.Tx, ctx context.Context, title string) error {
	var id int
	err := tx.QueryRow(ctx, `
		SELECT id FROM equipments WHERE title = $1 FOR UPDATE
	`, title).Scan(&id)

	if err == nil {
		err = fmt.Errorf("equipment with title '%s' already exists", title)
		return err
	}

	if err != pgx.ErrNoRows {

		return err
	}

	return nil
}
