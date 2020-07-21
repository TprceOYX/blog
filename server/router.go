package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(e *gin.Engine) {
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
