package main

import (
	"github.com/Juniordang/study-focus-api/cmd/config"
	"github.com/Juniordang/study-focus-api/internal/data/conf"
)

var logger *config.Logger

func main() {
	logger = config.Newlogger("main")

	err := conf.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}
}
