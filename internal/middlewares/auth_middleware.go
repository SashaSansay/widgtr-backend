package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"
	"widgtr-backend/internal/constants"
	"widgtr-backend/internal/controllers"
	"widgtr-backend/internal/models/user_model"
	"widgtr-backend/internal/utils/reponse_handlers"
)

var AuthMiddleware *jwt.GinJWTMiddleware

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func setupAuthorizationMiddleware(r *gin.Engine) {
	secretKey := os.Getenv("SECRET_KEY")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "widgtr",
		Key:         []byte(secretKey),
		Timeout:     time.Hour * 24 * 31,
		MaxRefresh:  time.Hour * 24 * 31,
		IdentityKey: constants.AuthIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			v, ok := data.(*user_model.User)
			if ok {
				claims := jwt.MapClaims{
					constants.AuthIdentityKey: v.ID,
				}
				return claims
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, _ := primitive.ObjectIDFromHex(claims[constants.AuthIdentityKey].(string))
			return id
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			reponse_handlers.Handle200(c, gin.H{
				"token":  token,
				"expire": expire.Unix(),
			})
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Email
			password := loginVals.Password

			user, err := user_model.GetByEmailFromDB(email)

			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			if user_model.CheckPassword(user.Password, password) {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			id, ok := data.(primitive.ObjectID)
			return ok && !id.IsZero()
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			reponse_handlers.Handle401(c, message)
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	auth := r.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/register", controllers.RegisterHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			user, _ := c.Get(constants.AuthIdentityKey)
			reponse_handlers.Handle200(c, user.(*user_model.User))
		})
	}

	AuthMiddleware = authMiddleware
}

func MetaUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get(constants.AuthIdentityKey)
		if !exists && userId.(primitive.ObjectID).IsZero() {
			reponse_handlers.Handle401(c, "unathorized")
			return
		}

		user, _ := user_model.GetFromDB(userId.(primitive.ObjectID).Hex())
		c.Set(constants.MetaUserKey, user)

		c.Next()
	}
}
