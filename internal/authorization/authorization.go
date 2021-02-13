package authorization

import (
	"context"

	"github.com/duyquang6/go-rbac-practice/internal/authorization/database"
	"github.com/duyquang6/go-rbac-practice/internal/authorization/model"
	"github.com/duyquang6/go-rbac-practice/pkg/dto"
	"github.com/duyquang6/go-rbac-practice/pkg/logging"
)

type AuthorizationService interface {
	CreateRole(ctx context.Context, role dto.CreateRoleRequest) error
	BindingUserRole(ctx context.Context, role_id int64) error
	CreatePolicy(ctx context.Context, policy dto.CreatePolicyRequest) error
	AppendPermissionPolicy(ctx context.Context, req dto.AppendPermissionPolicyRequest) error
	BindingPolicyRole(ctx context.Context, req dto.BindingPolicyRoleRequest) error
}

type authorizationService struct {
	authorizationRepo database.AuthorizationRepository
}

func NewAuthorizationService(authorizationRepo database.AuthorizationRepository) *authorizationService {
	return &authorizationService{authorizationRepo: authorizationRepo}
}

func (s *authorizationService) CreateRole(ctx context.Context, role dto.CreateRoleRequest) error {
	logger := logging.FromContext(ctx).Named("CreateRole")
	roleModel := model.Role{
		Name:        role.Name,
		Description: role.Description,
	}
	if err := s.authorizationRepo.CreateRole(ctx, roleModel); err != nil {
		logger.Error("create role failed: %v", err)
		return err
	}
	return nil
}

func (s *authorizationService) BindingUserRole(ctx context.Context, role_id int64) error {
	_ = logging.FromContext(ctx).Named("BindingUserRole")
	return nil
}

func (s *authorizationService) CreatePolicy(ctx context.Context, policy dto.CreatePolicyRequest) error {
	logger := logging.FromContext(ctx).Named("CreatePolicy")
	policyModel := model.Policy{
		Name:        policy.Name,
		Description: policy.Description,
	}
	if err := s.authorizationRepo.CreatePolicy(ctx, policyModel); err != nil {
		logger.Error("create policy failed: %v", err)
		return err
	}
	return nil
}

func (s *authorizationService) AppendPermissionPolicy(ctx context.Context, req dto.AppendPermissionPolicyRequest) error {
	logger := logging.FromContext(ctx).Named("AppendPermissionPolicy")
	var permissions []model.Permission

	for _, id := range req.PermissionIDs {
		permissions = append(permissions, model.Permission{ID: id})
	}

	if err := s.authorizationRepo.AppendPermissionPolicy(ctx, req.PolicyID, permissions); err != nil {
		logger.Error("append permission policy failed: %v", err)
		return err
	}
	return nil
}

func (s *authorizationService) BindingPolicyRole(ctx context.Context, req dto.BindingPolicyRoleRequest) error {
	logger := logging.FromContext(ctx).Named("BindingPolicyRole")
	var policies []model.Policy

	for _, id := range req.PolicyIDs {
		policies = append(policies, model.Policy{ID: id})
	}

	if err := s.authorizationRepo.BindingPolicyRole(ctx, req.RoleID, policies); err != nil {
		logger.Error("binding policy role failed: %v", err)
		return err
	}
	return nil
}
