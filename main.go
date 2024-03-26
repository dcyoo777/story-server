package main

import (
	"example/router"
	"github.com/gin-contrib/cors"
)

func main() {
	r := router.SetupRouter()
	r.Use(cors.Default())
	r.Run(":8080")
}
