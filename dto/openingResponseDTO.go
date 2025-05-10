package dto

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message":    message,
		"statusCode": code,
	})
}

func SendSuccessResponse(ctx *gin.Context, op string, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"message":    fmt.Sprintf("Operation successful: %s", op),
		"data":       data,
	})
}
