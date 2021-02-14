package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/duyquang6/go-rbac-practice/internal/controller"
	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/gin-gonic/gin"
)

func (s *Controller) HandleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		session := controller.SessionFromContext(ctx)
		if session == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		req := dto.LoginRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err = s.authService.Login(ctx, session, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
