package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserInfo struct {
	TgUS string
	TgID string
}

type NFTs_Collection struct {
	Name              string
	Image_url         string
	Base_price        float64
	Crypto_connection string
	Max_supply        int
	Mint              int
}

type Coins struct {
	Coin_name  string
	Coin_id    string
	Coin_image string
	Price      int
	Change_24h float64
}

type NFTSsOwners struct {
	Nft_id                        int
	User_id                       int
	Quantity                      int
	Price                         float64
	Rarity                        string
	Upgrades_image_url            string
	Total_spent                   float64
	Upgrades_background_image_url string
	Is_busters                    bool
}

type Subscribe struct {
	Time  string
	Price int
}

func InsertRowUsers(ctx context.Context, connection *pgx.Conn, data UserInfo) error {
	sqlQuery := `
	INSERT INTO users (telegram_id, telegram_username, subscribe, balance)
	VALUES ($1,$2,DEFAULT, DEFAULT)
	`
	_, err := connection.Exec(ctx, sqlQuery, data.TgID, data.TgUS)

	if err != nil {
		return err
	}
	return nil
}

func InsertRowNFTs(ctx context.Context, connection *pgx.Conn, data NFTs_Collection) error {
	sqlQuery := `
	INSERT INTO nfts (name, image_url, base_price, crypto_connection, max_supply, created_at, mint)
	VALUES ($1,$2,$3,$4,$5,DEFAULT,$6)
	`
	_, err := connection.Exec(ctx, sqlQuery, data.Name, data.Image_url, data.Base_price, data.Crypto_connection, data.Max_supply, data.Max_supply)
	if err != nil {
		return err
	}
	return nil
}

func InsertRowCoins(ctx context.Context, connection *pgx.Conn, data Coins) error {
	sqlQuery := `
	INSERT INTO coins (coin_name, coin_id, coin_image, price, change_24h)
	VALUES ($1,$2,$3,$4,$5)
	`

	_, err := connection.Exec(ctx, sqlQuery, data.Coin_name, data.Coin_id, data.Coin_image, data.Price, data.Change_24h)
	if err != nil {
		return err
	}
	return nil
}

func InsertRowNFTsOwners(ctx context.Context, connection *pgx.Conn, data NFTSsOwners) error {
	sqlQuery := `
	INSERT INTO nfts_owners (nft_id, user_id, quantity, total_spent, price, last_update, rarity, upgrades_image_url, upgrades_background_image_url, is_busters)
	VALUES ($1,$2,$3,$4,$5,DEFAULT,$6,$7,$8,DEFAULT)
	`
	_, err := connection.Exec(ctx, sqlQuery, data.Nft_id, data.User_id, data.Quantity, data.Total_spent, data.Price, data.Rarity, data.Upgrades_image_url, data.Upgrades_background_image_url)
	if err != nil {
		return err
	}
	return nil
}

func InsertRowSubscribe(ctx context.Context, connection *pgx.Conn, data Subscribe) error {
	sqlQuery := `
	INSERT INTO subscribe (time, price)
	VALUES ($1,$2)
	`
	_, err := connection.Exec(ctx, sqlQuery, data.Time, data.Price)
	if err != nil {
		return err
	}
	return nil
}
