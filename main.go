package main

import (
	"github.com/arthurlopesr/go-opportunities/config"
	"github.com/arthurlopesr/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorFormated("Config initialization failed: %v", err)
		return
	}
	router.Initialize()
}
