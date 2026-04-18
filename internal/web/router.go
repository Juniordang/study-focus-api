package web

import (
	"github.com/Juniordang/study-focus-api/cmd/config"
	"github.com/gin-gonic/gin"
)

func Initialize() error {
	logger := config.Newlogger("router")
	router := gin.Default()

	InitializeRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logger.Errf("error on start server: %v", err)
		return err
	}

	return nil
}
