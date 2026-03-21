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
	if err = database.CreateTable(ctx, conn); err != nil {
		fmt.Println(err)
	}
	err = database.CreateTableNft(ctx, conn)
	if err != nil {
		fmt.Println(err)
	}

}
