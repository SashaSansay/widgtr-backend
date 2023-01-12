package controllers

import (
	"github.com/gin-gonic/gin"
	"widgtr-backend/internal/constants"
	"widgtr-backend/internal/utils/reponse_handlers"
)

func MeIndexHandler(c *gin.Context) {
	user, _ := c.Get(constants.MetaUserKey)

	reponse_handlers.Handle200(c, user)
}
