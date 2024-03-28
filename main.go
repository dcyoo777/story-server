package main

import (
	"example/request"
	"example/router"
	"example/service"
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

	service.StoryCommonReq = service.StoryCommonRequests{
		CommonRequests: request.CommonRequests{
			Table:      "story",
			PrimaryKey: "story_id",
			DatasourceName: fmt.Sprintf(
				"%s:%s@tcp(%s:3306)/%s",
				os.Getenv("DB_USERNAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("SERVER_DOMAIN"),
				os.Getenv("DB_NAME"),
			),
		},
	}

	r := router.SetupRouter()
	r.Run(":80")
}
