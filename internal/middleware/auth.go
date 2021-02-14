package middleware

import (
	"net/http"

	"github.com/duyquang6/go-rbac-practice/internal/auth"
	"github.com/duyquang6/go-rbac-practice/internal/controller"
	"github.com/duyquang6/go-rbac-practice/internal/user"
	"github.com/gin-gonic/gin"
)

func AuthSession(authService auth.AuthService, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		session := controller.SessionFromContext(ctx)
		if session == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sessionData, err := authService.GetSessionData(ctx, session)
		// Need verify expiration
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := userService.GetByUsername(ctx, sessionData.Username)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx = controller.WithPermissionMapping(ctx, sessionData.Permission)
		ctx = controller.WithUser(ctx, user)
		c.Request = c.Request.Clone(ctx)

		c.Next()
	}
}
