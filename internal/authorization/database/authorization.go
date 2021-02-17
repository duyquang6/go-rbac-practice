package database

import (
	"context"

	authorizedModel "github.com/duyquang6/go-rbac-practice/internal/authorization/model"

	"github.com/duyquang6/go-rbac-practice/internal/database"
)

type AuthorizationRepository interface {
	CreateRole(ctx context.Context, role authorizedModel.Role) error
	BindingRole(ctx context.Context, role_id int64) error
	CreatePolicy(ctx context.Context, policy authorizedModel.Policy) error
	AppendPermissionPolicy(ctx context.Context, policyID int64, permissions []authorizedModel.Permission) error
	BindingPolicyRole(ctx context.Context, roleID int64, policies []authorizedModel.Policy) error
	GetPermissionByUserID(ctx context.Context, userID int64) ([]authorizedModel.Permission, error)
}

type authorizationDB struct {
	db *database.DB
}

func New(db *database.DB) *authorizationDB {
	return &authorizationDB{
		db: db,
	}
}

func (s *authorizationDB) CreateRole(ctx context.Context, role authorizedModel.Role) error {
	res := s.db.Pool.Create(&role)
	return res.Error
}

func (s *authorizationDB) BindingRole(ctx context.Context, role_id int64) error {
	return nil
}

func (s *authorizationDB) CreatePolicy(ctx context.Context, policy authorizedModel.Policy) error {
	res := s.db.Pool.Create(&policy)
	return res.Error
}

func (s *authorizationDB) AppendPermissionPolicy(ctx context.Context, policyID int64, permissions []authorizedModel.Permission) error {
	return s.db.Pool.Model(&authorizedModel.Policy{ID: policyID}).Association("Permissions").Append(permissions)
}

func (s *authorizationDB) BindingPolicyRole(ctx context.Context, roleID int64, policies []authorizedModel.Policy) error {
	return s.db.Pool.Model(&authorizedModel.Role{ID: roleID}).Association("Policies").Append(policies)
}

func (s *authorizationDB) GetPermissionByUserID(ctx context.Context, userID int64) ([]authorizedModel.Permission, error) {
	var (
		permissions []authorizedModel.Permission
	)
	err := s.db.Pool.Raw(`SELECT p.* FROM permissions p
							INNER JOIN permission_policies pp ON p.id = pp.permission_id
							INNER JOIN policy_roles pr on pp.policy_id = pr.policy_id
							INNER JOIN user_roles ur on ur.role_id = pr.role_id
							WHERE user_id=?`, userID).Scan(&permissions).Error
	return permissions, err
}
