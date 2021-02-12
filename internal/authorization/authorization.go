package authorization

import (
	"context"

	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
)

type AuthorizationService interface {
	CreateRole(ctx context.Context, role dto.CreateRoleRequest) error
	BindingRole(ctx context.Context, role_id int64) error
	CreatePolicy(ctx context.Context, policy dto.CreateRoleRequest) error
}

type authService struct{}

func NewAuthorizationService() *authService {
	return &authService{}
}

func (s *authService) CreateRole(ctx context.Context, role dto.CreateRoleRequest) error {
	var err error
	logger := logging.FromContext(ctx).Named("CreateRole")
	logger.Error("create role failed: %v", err)
	return nil
}

func (s *authService) BindingRole(ctx context.Context, role_id int64) error {
	return nil

}

func (s *authService) CreatePolicy(ctx context.Context, policy dto.CreateRoleRequest) error {
	return nil

}
