package router

import (
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group(constant.APIBasePath)
	{
		v1.GET(constant.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
			})
		})
		v1.POST(constant.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": " post Hello, World!",
			})
		})
		v1.DELETE(constant.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Delete Hello, World!",
			})
		})
		v1.PUT(constant.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Hello, World!",
			})
		})
		v1.GET(constant.APIOpeningsPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "find all Hello, World!",
			})
		})
	}
}
