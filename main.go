package main

import (
	"example/mysql"
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
		"%s:%s@tcp(%s:3306)/",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SERVER_DOMAIN"),
	)

	_, err = mysql.InitDB(dataSourceName, os.Getenv("DB_NAME"))
	if err != nil {
		return
	}

	r := router.SetupRouter(dataSourceName + os.Getenv("DB_NAME"))
	r.Run(":" + os.Getenv("PORT"))
}
