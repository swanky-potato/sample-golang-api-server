package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func Enviroment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, os.Environ())
	}
}
