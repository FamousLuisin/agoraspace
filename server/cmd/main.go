package main

import (
	"log"

	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/FamousLuisin/agoraspace/internal/routes"
	"github.com/FamousLuisin/agoraspace/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	connect, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer connect.Db.Close()

	if err = connect.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = db.Migrations(connect.Db); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, connect)

	port, err := utils.GetEnv("SERVER_HOST")

	if err != nil {
		log.Fatal(err)
	}

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}