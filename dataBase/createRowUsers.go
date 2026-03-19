package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserInfo struct {
	TgUS string
	TgID string
}

func InsertRowUsers(ctx context.Context, connection *pgx.Conn, data UserInfo) error {
	sqlQuery := `
	INSERT INTO users (telegram_id, telegram_username, subscribe, balance)
	VALUES ($1,$2,DEFAULT, DEFAULT)
	`
	_, err := connection.Exec(ctx, sqlQuery, data.TgUS, data.TgID)

	if err != nil {
		return err
	}
	return nil
}
