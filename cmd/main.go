package main

import (
	"log"
	"os"

	"auth/api"
	"auth/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	envErr := config.CheckENV()

	if envErr != nil {
		log.Fatalln(envErr)
	}

	dsn := os.Getenv("DSN")

	dbConn, db, err := config.OpenDB(dsn)

	if err != nil {
		log.Fatalln(err)
	}

	config.CheckForMigration(db)

	app := config.SetupConfig(dbConn, db)

	r := api.SetupRouter(&app.API)
	r.Run(":8081")
}
