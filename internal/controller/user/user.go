package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/duyquang6/go-rbac-practice/internal/controller"
	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/duyquang6/go-rbac-practice/pkg/rbac"
	"github.com/gin-gonic/gin"
)

func (s *Controller) HandleCreateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		permissionMapping := controller.PermissionMappingFromContext(ctx)
		if !rbac.IsPermit(permissionMapping, rbac.User, rbac.Create) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		req := dto.CreateUserRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err = s.userService.CreateUser(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}

func (s *Controller) HandleBindingRoleUser() func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		req := dto.BindingRoleUserRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		req.UserID, err = strconv.ParseInt((c.Param("id")), 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = s.userService.BindingRoleUser(ctx, req); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
