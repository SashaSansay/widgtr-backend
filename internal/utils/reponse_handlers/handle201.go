package reponse_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle201(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "data": data, "success": true})
}
