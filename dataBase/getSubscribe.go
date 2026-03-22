package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type BuySubscribe struct {
	TelegramID      string
	SubscribeAmount string
}

func GetSubscribe(ctx context.Context, connection *pgx.Conn, data BuySubscribe) (bool, error) {
	var subscribePrice int
	var userBalance float64
	sqlQueryGetPrice := `
	SELECT price
	FROM subscribe
	WHERE time = $1
	`
	sqlQueryGetUserBalance := `
	SELECT balance
	FROM users
	WHERE telegram_id = $1
	`
	err := connection.QueryRow(ctx, sqlQueryGetPrice, data.SubscribeAmount).Scan(&subscribePrice)
	if err != nil {
		return false, err
	}
	err = connection.QueryRow(ctx, sqlQueryGetUserBalance, data.TelegramID).Scan(&userBalance)
	if err != nil {
		return false, err
	}
	if userBalance >= subscribePrice && isSubscribe == false {
		transaction, err := connection.Begin(ctx)
		if err != nil {
			return false, err
		}
		defer transaction.RollBack(ctx)
		sqlQueryUpdateUserBalnce := `
		UPDATE users
		SET balance = balance - $1
		WHERE telegram_id = $2
		`
		_, err = transaction.Exec(ctx, sqlQueryUpdateUserBalnce, subscribePrice, data.TelegramID)
		if err != nil {
			return false, err
		}
	}

}
