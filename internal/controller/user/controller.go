package user

import (
	"github.com/duyquang6/go-rbac-practice/internal/user"
)

type Controller struct {
	userService user.UserService
}

// New creates a new user controller.
func New(userService user.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}
