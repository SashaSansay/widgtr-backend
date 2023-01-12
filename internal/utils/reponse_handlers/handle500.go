package reponse_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle500(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "error": err, "success": false})
}
