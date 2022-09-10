package api

// import (
// 	"github.com/gin-gonic/gin"
// )

// // CORS (Cross-Origin Resource Sharing)
// // CORS middleware is always very specific to use case,
// // so adding very minimal placeholder here
// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "600")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
// 		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 		} else {
// 			c.Next()
// 		}
// 	}
// }
