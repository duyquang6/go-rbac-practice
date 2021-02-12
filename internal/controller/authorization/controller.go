package authorization

import (
	"github.com/duyquang6/go-rbac-practice/internal/authorization"
)

type Controller struct {
	authService authorization.AuthorizationService
}

// New creates a new authorization controller.
func New(authService authorization.AuthorizationService) *Controller {
	return &Controller{
		authService: authService,
	}
}
