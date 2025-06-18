package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/FamousLuisin/agoraspace/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connection() (*pgxpool.Pool,error){
	database, _database := utils.GetEnv("DB_NAME")
	user, _user := utils.GetEnv("DB_USER")
	password, _password := utils.GetEnv("DB_PASSWORD")
	host, _host := utils.GetEnv("DB_HOST")
	port, _port := utils.GetEnv("DB_PORT")

	if err := errors.Join(_database, _user, _password, _host, _port); err != nil  {
		return nil, fmt.Errorf("error getting database environment variables:\n%s", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to parse pool config: %w", err)
    }

    pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
        return nil, fmt.Errorf("failed to create pool: %w", err)
    }

	return pool, nil
}