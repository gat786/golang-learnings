package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupApiRoute(rg *gin.RouterGroup) {
	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong from api group",
		})
	})
}
