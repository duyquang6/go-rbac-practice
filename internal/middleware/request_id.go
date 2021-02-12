package middleware

import (
	"net/http"

	"github.com/duyquang6/go-rbac-practice/internal/controller"
	"github.com/duyquang6/go-rbac-practice/pkg/customuid"
	"github.com/gin-gonic/gin"
)

const (
	requestIDHeader = "X-Request-ID"
)

func PopulateRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		ctx := c.Request.Context()

		if _, ok := c.Request.Header[http.CanonicalHeaderKey(requestIDHeader)]; ok {
			requestID = c.Request.Header[http.CanonicalHeaderKey(requestIDHeader)][0]
		} else {
			requestID = customuid.GetUniqueID()
		}

		ctx = controller.WithRequestID(ctx, requestID)
		c.Request = c.Request.Clone(ctx)

		c.Next()
	}
}
