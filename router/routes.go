package router

import (
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/arthurlopesr/go-opportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitHandler()
	v1 := router.Group(constant.APIBasePath)
	{
		v1.GET(constant.APIOpeningPath, handler.FindOpeningHandler)
		v1.POST(constant.APIOpeningPath, handler.CreateOpeningHandler)
		v1.DELETE(constant.APIOpeningPath, handler.DeleteOpeningHandler)
		v1.PUT(constant.APIOpeningPath, handler.UpdateOpeningHandler)
		v1.GET(constant.APIOpeningsPath, handler.FindAllOpeningHandler)
	}
}
