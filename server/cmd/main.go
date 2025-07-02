package main

import (
	"fmt"

	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println(err)
	}

	connect, err := db.Connection()
	if err != nil {
		panic(err)
	}
	defer connect.Db.Close()

	if err = connect.Db.Ping(); err != nil {
		panic(err)
	}

	if err = db.Migrations(connect.Db); err != nil {
		panic(err)
	}
}