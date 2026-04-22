package student

import (
	"github.com/Juniordang/study-focus-api/internal/data/schema"
	"github.com/Juniordang/study-focus-api/internal/web/handlers"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"nome" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"senha" binding:"required,min=8"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.SendError(c, 400, err.Error())
		return
	}

	var user schema.User

	handlers.DB.First(&user, "email = ?", input.Email)
	if user.Model.ID != 0 {
		handlers.SendError(c, 409, "Student already exists")
		return
	}

	user = schema.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := handlers.DB.WithContext(c.Request.Context()).Create(&user); err.Error != nil {
		handlers.SendError(c, 400, "Not create User")
		return
	}

	handlers.SendSuccess(c, 201, user)
}
