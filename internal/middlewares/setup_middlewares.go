package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetupMiddlewares(r *gin.Engine) {
	setupCorsMiddleware(r)
	setupAuthorizationMiddleware(r)
}
