package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		telegram_id VARCHAR(70) UNIQUE NOT NULL,
		telegram_username VARCHAR(70) UNIQUE NOT NULL,
		subscribe BOOLEAN DEFAULT FALSE,
		balance DECIMAL(20,2) DEFAULT 0.00,
		subscribe_end_date TIMESTAMP
	);
	`

	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

func CreateTableNft(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS nfts(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	image_url TEXT NOT NULL,
	rarity VARCHAR(100) NOT NULL,
	base_price NUMERIC N0T NULL,
	crypto_connection VARCHAR(100) NOT NULL,
	max_supply INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	mint INTEGER NOT NULL,
	);
	`

	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}
