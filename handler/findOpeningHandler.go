package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "find opening",
	})
}
