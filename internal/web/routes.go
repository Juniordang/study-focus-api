package web

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		users := v1.Group("/users")
		{
			users.POST("/")
		}
	}
}
