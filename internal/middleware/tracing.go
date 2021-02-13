package middleware

import "github.com/gin-gonic/gin"

// PopulateTracing
func PopulateTracing() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
