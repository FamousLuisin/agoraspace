package db

import (
	"errors"
	"fmt"

	"github.com/FamousLuisin/agoraspace/internal/utils"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Db *sqlx.DB 
}

func Connection() (*Database, error){
	database, _database := utils.GetEnv("DB_NAME")
	user, _user := utils.GetEnv("DB_USER")
	password, _password := utils.GetEnv("DB_PASSWORD")
	host, _host := utils.GetEnv("DB_HOST")
	port, _port := utils.GetEnv("DB_PORT")

	if err := errors.Join(_database, _user, _password, _host, _port); err != nil  {
		return nil, fmt.Errorf("error getting database environment variables:\n%s", err)
	}
	
	config := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, database, host, port)

	db, _db := sqlx.Connect("pgx", config)
    
	if _db != nil {
		return nil, fmt.Errorf("error connecting to database -> %s", _db)
	}

	return &Database{
		Db: db,
	}, nil
}

func Migrations(db *sqlx.DB) (error){
	driver, _driver := postgres.WithInstance(db.DB, &postgres.Config{})
	if _driver != nil {
		return fmt.Errorf("error getting driver instance -> %s", _driver)
	}

	migration, _migration := migrate.NewWithDatabaseInstance("file://./migrations", "agoraspace", driver)
	if _migration != nil {
		return fmt.Errorf("error getting the migrate instance -> %s", _migration)
	}

	_up := migration.Up();
	if _up != nil && _up != migrate.ErrNoChange {
		return fmt.Errorf("error running migrate -> %s", _up)
	}

	if _up == migrate.ErrNoChange {
		fmt.Println("no migration changes to apply.")
	}

	return nil
}