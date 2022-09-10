package api

import (
	"context"
	"crud-with-auth/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

type Api struct {
	r     *gin.Engine
	db    db.ProviderDB
	Cache *redis.Client
}

var RedisContext = context.Background()

func (api *Api) Start() {
	api.r.GET("/", api.HomeHandler)

	api.Auth()
	api.Articles()
	api.Leaderboard()
	api.OAuth()

	api.r.Run(":4000")
}

func (api Api) HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
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
