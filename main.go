package main

import (
	"context"
	"fmt"
	database "main/dataBase"
	dbconnection "main/dataBase/DBconnection"
)

func main() {
	ctx := context.Background()
	conn, err := dbconnection.ConnectionToBD(ctx)
	if err != nil {
		fmt.Println("При подключении произошла ошибка")
	} else {
		fmt.Println("Подключение установленно")
	}

	t := database.BuySubscribe{
		TelegramID:      "123123",
		SubscribeAmount: "365",
	}
	p, err := database.GetSubscribe(ctx, conn, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
