package middlewares

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// NewGzip return gzip middleware compression
func NewGzip() gin.HandlerFunc {
	return gzip.Gzip(gzip.BestSpeed)
}