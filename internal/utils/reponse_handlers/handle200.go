package reponse_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "success": true})
}
