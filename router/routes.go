package router

import (
	"github.com/arthurlopesr/go-opportunities/router/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group(constants.APIBasePath)
	{
		v1.GET(constants.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
			})
		})
		v1.POST(constants.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": " post Hello, World!",
			})
		})
		v1.DELETE(constants.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Delete Hello, World!",
			})
		})
		v1.PUT(constants.APIOpeningPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Hello, World!",
			})
		})
		v1.GET(constants.APIOpeningsPath, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "find all Hello, World!",
			})
		})
	}
}
