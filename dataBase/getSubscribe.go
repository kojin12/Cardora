package database

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type BuySubscribe struct {
	TelegramID      string
	SubscribeAmount string
}

type UserData struct {
	ID           int
	Balance      float64
	SubscribeEnd time.Time
}

func GetSubscribe(ctx context.Context, connection *pgx.Conn, data BuySubscribe) (bool, error) {
	tx, err := connection.Begin(ctx)
	if err != nil {
		return false, err
	}
	defer tx.Rollback(ctx)

	var user UserData
	var price float64

	if err = tx.QueryRow(ctx, `
        SELECT id, balance, subscribe_end_date 
        FROM users 
        WHERE telegram_id=$1 FOR UPDATE
    `, data.TelegramID).Scan(&user.ID, &user.Balance, &user.SubscribeEnd); err != nil {
		return false, err
	}

	if err = tx.QueryRow(ctx, "SELECT price FROM subscribe WHERE time=$1", data.SubscribeAmount).Scan(&price); err != nil {
		return false, err
	}

	if user.Balance < price {
		return false, errors.New("Недостаточно средств")
	}

	now := time.Now().UTC()
	if now.Before(user.SubscribeEnd) {
		return false, errors.New("Ваша подписка ещё действительна")
	}

	days, err := strconv.Atoi(data.SubscribeAmount)
	if err != nil {
		return false, errors.New("Неверное значение подписки")
	}

	var newSubscribeEnd time.Time
	if now.After(user.SubscribeEnd) {
		newSubscribeEnd = now.AddDate(0, 0, days)
	} else {
		newSubscribeEnd = user.SubscribeEnd.AddDate(0, 0, days)
	}

	_, err = tx.Exec(ctx, `
        UPDATE users 
        SET balance = balance - $1,
            subscribe_end_date = $2
        WHERE id = $3
    `, price, newSubscribeEnd, user.ID)
	if err != nil {
		return false, err
	}

	if err := tx.Commit(ctx); err != nil {
		return false, err
	}

	return true, nil
}
