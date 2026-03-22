package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Таблица юзеров
func CreateTableUsers(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		telegram_id VARCHAR(70) UNIQUE NOT NULL,
		telegram_username VARCHAR(70) UNIQUE NOT NULL,
		balance DECIMAL(20,2) DEFAULT 0.0,
		subscribe_end_date TIMESTAMP,
		wallets JSONB DEFAULT '[]'
	);
	`

	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

// Таблица коллекций нфт
func CreateTableNft(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS nfts(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		image_url TEXT NOT NULL,
		background_url TEXT DEFAULT NULL,
		base_price NUMERIC NOT NULL,
		crypto_connection INTEGER REFERENCES coins(id),
		max_supply INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		mint INTEGER NOT NULL,
		multiplier NUMERIC(20,8) DEFAULT 0.0
	);
	`

	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

// Таблица держателей нфт
func CreateTableNftOwners(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS nfts_owners(
		id SERIAL PRIMARY KEY,
		nft_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		quantity INTEGER NOT NULL DEFAULT 1,
		total_spent NUMERIC(20,8) NOT NULL DEFAULT 0.0,
		price NUMERIC(20,8) NOT NULL,
		last_update TIMESTAMP DEFAULT NOW(),
		rarity VARCHAR(50) DEFAULT NULL,
		upgrades_image_url TEXT DEFAULT NULL,
		upgrades_background_image_url TEXT DEFAULT NULL,
		is_busters BOOLEAN DEFAULT FALSE
	);
	`

	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
	return nil
}

// Таблица монет
func CreateTableCoin(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS coins(
		id SERIAL PRIMARY KEY,
		coin_name VARCHAR(70) NOT NULL,
		coin_id VARCHAR(10) NOT NULL,
		coin_image TEXT DEFAULT NULL,
		price NUMERIC(25,10) DEFAULT NULL,
		change_24h NUMERIC(20,8) DEFAULT NULL
	);
	`
	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	return nil
}

// Таблица подписки
func CreateTableSubscribe(ctx context.Context, connection *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS subscribe(
		id SERIAL PRIMARY KEY,
		time VARCHAR(100) UNIQUE,
		price INTEGER DEFAULT 0
	);
	`
	_, err := connection.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	return nil
}
