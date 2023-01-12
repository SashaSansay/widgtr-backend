package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"widgtr-backend/internal/controllers"
	"widgtr-backend/internal/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.
		Group("/api/v1").
		GET("/api/v1/widgets", controllers.FindAllWidgetsHandler).
		POST("/api/v1/widgets", controllers.CreateWidgetHandler)

	r.Use(middlewares.AuthMiddleware.MiddlewareFunc())
	{
		r.Use(middlewares.MetaUserMiddleware())
		{
			r.
				Group("/api/v1/me").
				GET("/", controllers.MeIndexHandler)

			r.
				Group("/api/v1/pages").
				POST("/", controllers.CreatePageHandler)

			r.
				Group("/api/v1/widgets").
				POST("/", controllers.CreateWidgetHandler)
		}
	}
}
