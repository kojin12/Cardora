package dbconnection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectionToBD(ctx context.Context) (*pgx.Conn, error) {
	dbHOST := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")
	dbUSER := os.Getenv("DB_USER")
	dbPASSWORD := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUSER, dbPASSWORD, dbHOST, dbPORT, dbNAME)
	conn, err := pgx.Connect(ctx, "postgres://postgres:15Card47019_!DORA_77Null@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
