package database

import (
	"backend-balto/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDb(ctx context.Context, conf *models.DbConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", conf.Username, conf.Password,
		conf.Host, conf.Port, conf.Database)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Success connect Database : ", conf.Host)

	return conn, nil
}
