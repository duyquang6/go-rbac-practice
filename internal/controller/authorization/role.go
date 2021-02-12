package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/gin-gonic/gin"
)

func (s *Controller) HandleCreateRole() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		req := dto.CreateRoleRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err = s.authService.CreateRole(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
