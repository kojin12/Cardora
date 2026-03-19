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

	//5150505792, @Diozarm
	err = database.InsertRow(ctx, conn, "@Diozarm", "5150505792")
	if err != nil {
		fmt.Println("Ошибка, пользователь не добавлен:", err)
	}
	err = database.InsertRow(ctx, conn, "@emi4ka21", "1488398562")
	if err != nil {
		fmt.Println("Ошибка, пользователь не добавлен:", err)
	}
}
