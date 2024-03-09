package main

import (
	"example/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
