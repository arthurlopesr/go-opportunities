package router

import (
	"github.com/arthurlopesr/go-opportunities/constant"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(constant.APIPort)
}
