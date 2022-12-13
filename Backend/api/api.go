// @title Swagger Example API
// @version 1.0
// @description This is a description sample.
// @termsOfService http://swagger.io/terms/

// @contact.name Contact API Support
// @contact.url http://www.swagger.io/support
// @contact.email zkdltid.chan@gmail.com

// @host localhost:4000
// @BasePath /

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl http://localhost:4000/oauth/
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

package api

import (
	"context"
	"crud-with-auth/db"
	"log"
	"net/http"
	"os"
	"strings"

	docs "crud-with-auth/docs"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Api struct {
	r     *gin.Engine
	db    db.ProviderDB
	Cache *redis.Client
}

var RedisContext = context.Background()

func (api *Api) Start() {
	docs.SwaggerInfo.Title = "Auth Service API"
	api.r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.r.GET("/", api.HomeHandler)

	api.Auth()
	// api.Articles()
	api.User()
	api.Admin()
	api.Leaderboard()
	api.OAuth()
	api.r.Run(":4000")
}

func (api Api) HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func (api Api) AuthAdminMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	splittedToken := strings.Split(tokenString, " ")
	token := splittedToken[1]

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	var admin db.Admin
	if _, err := admin.ValidateAdminToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token invalid or expired",
		})
		return
	}

	c.Set("username", admin.Username)
	c.Set("admin_id", admin.ID)

	c.Next()
}

func (api Api) AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	splittedToken := strings.Split(tokenString, " ")
	token := splittedToken[1]

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	var user db.User
	if _, err := user.ValidateToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token invalid or expired",
		})
		return
	}

	c.Set("username", user.Username)
	c.Set("user_id", user.ID)

	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "600")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func NewAPI(db db.ProviderDB) *Api {
	r := gin.Default()
	r.Use(CORSMiddleware())

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		// Addr: "redis:6379",
	})

	if err := rdb.Ping(RedisContext).Err(); err != nil {
		panic("failed ping redis")
	} else {
		log.Println("redis is connected")
	}

	return &Api{r: r, db: db, Cache: rdb}
}

var ProviderAPI = wire.NewSet(NewAPI)
