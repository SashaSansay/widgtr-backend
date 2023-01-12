package reponse_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle401(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": message, "success": false})
}
