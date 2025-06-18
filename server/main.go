package main

import (
	"fmt"

	"github.com/FamousLuisin/agoraspace/db"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	connect, err := db.Connection()
	if err != nil {
		panic(err)
	}
	defer connect.Close()

	if err = connect.Ping(); err != nil {
		panic(err)
	}

	if err = db.Migrations(connect); err != nil {
		panic(err)
	}
}