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

	gingonic := gin.Default()

	routes.InitRoutes(&gingonic.RouterGroup)

	port, err := utils.GetEnv("SERVER_HOST")

	if err != nil {
		log.Fatal(err)
	}

	if err := gingonic.Run(port); err != nil {
		log.Fatal(err)
	}
}