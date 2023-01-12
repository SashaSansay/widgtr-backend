package reponse_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle400(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": err.Error(), "success": false})
}
