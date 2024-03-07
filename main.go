package main

import (
	"next-month/service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.Get(app)
	app.Run()
}
