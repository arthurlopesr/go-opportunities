package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
