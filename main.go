package main

import (
	"SampleGoGin/database"
	"SampleGoGin/middlewares"
	. "SampleGoGin/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()
	router := gin.Default()

	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger())

	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DB()
	}()

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message":  "ping",
	// 		"message2": "ping2",
	// 	})
	// })
	// router.POST("/ping/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })

	router.Run(":8000")
}
