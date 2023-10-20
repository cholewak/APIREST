package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
}
