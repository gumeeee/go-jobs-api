package main

import (
	"github.com/gumeeee/go-jobs-api/config"
	"github.com/gumeeee/go-jobs-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
