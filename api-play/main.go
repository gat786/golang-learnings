package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type hello struct {
	Message string `json:"message"`
	Content string `json:"content"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, hello{
			Message: "Hello World",
			Content: "This is some random content that I want to test",
		})
	})

	waveMessage := hello{
		Message: "Hello World",
		Content: "This is message which comes from a struct which is created outside of gin GET.",
	}

	router.GET("/wave", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &waveMessage)
	})

	apiRouter := router.Group("/api")
	setupApiRoute(apiRouter)

	router.Run(":8080")
}
