package main

import (
	"example/router"
	_ "example/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SERVER_DOMAIN"),
		os.Getenv("DB_NAME"),
	)

	r := router.SetupRouter(dataSourceName)
	r.Run(":" + os.Getenv("PORT"))
}
