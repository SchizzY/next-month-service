package routes

import (
	"database/sql"
	"next-month/service/handler"

	"github.com/gin-gonic/gin"
)

func Get(r *gin.Engine, db *sql.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", func(c *gin.Context) {
		handler.CreateTest(db)

		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/test-drop", func(c *gin.Context) {
		handler.ClearTest(db)

		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
