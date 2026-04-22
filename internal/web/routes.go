package web

import (
	"github.com/Juniordang/study-focus-api/internal/data/conf"
	"github.com/Juniordang/study-focus-api/internal/web/handlers"
	"github.com/Juniordang/study-focus-api/internal/web/handlers/student"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	handlers.DB = conf.GetSqlite()

	v1 := r.Group("/api/v1")

	{
		users := v1.Group("/users")
		{
			users.POST("/", student.CreateUser)
		}
	}
}
