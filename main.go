package main

import (
	"example/router"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	r := router.SetupRouter()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			MaxAge:       12 * time.Hour,
		}))
	r.Run(":8080")
}
