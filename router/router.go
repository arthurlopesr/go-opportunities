package router

import (
	"github.com/arthurlopesr/go-opportunities/router/constants"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(constants.APIPort)
}
