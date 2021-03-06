package authorization

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
		if err = s.authorizationService.CreateRole(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

func (s *Controller) HandleBindingPolicyRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		req := dto.BindingPolicyRoleRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.RoleID, err = strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = s.authorizationService.BindingPolicyRole(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
