package main

import (
	"gobooks/config"
	"gobooks/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errf("Config initialization error: %v", err)
		panic(err)
	}
	router.Initialize()
}
