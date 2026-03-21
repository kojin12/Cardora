package dbconnection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectionToBD(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, "postgres://postgres:15Card47019_!DORA_77Null@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
