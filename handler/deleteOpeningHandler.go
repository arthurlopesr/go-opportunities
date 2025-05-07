package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete opening",
	})
}
