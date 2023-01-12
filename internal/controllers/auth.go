package controllers

import (
	"github.com/gin-gonic/gin"
	"strings"
	"widgtr-backend/internal/models/user_model"
	"widgtr-backend/internal/utils/reponse_handlers"
)

type RegisterIn struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func RegisterHandler(c *gin.Context) {
	var input RegisterIn

	if err := c.ShouldBindJSON(&input); err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	hashedPassword, _ := user_model.HashPassword(input.Password)

	u := user_model.User{
		Email:     strings.TrimSpace(strings.ToLower(input.Email)),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  hashedPassword,
	}

	user, err := user_model.NewUser(u).InsertIntoDB()

	if err != nil {
		reponse_handlers.Handle400(c, err)
		return
	}

	reponse_handlers.Handle201(c, user)
}
