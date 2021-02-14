package auth

import (
	"github.com/duyquang6/go-rbac-practice/internal/auth"
)

type Controller struct {
	authService auth.AuthService
}

// New creates a new authentication controller.
func New(authService auth.AuthService) *Controller {
	return &Controller{
		authService: authService,
	}
}
